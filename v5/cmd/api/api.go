package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"v5/internal/auth"
	"v5/internal/env"
	"v5/internal/mailer"
	"v5/internal/ratelimiter"
	"v5/internal/store"
	"v5/internal/store/cache"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"go.uber.org/zap"
)

type application struct {
	config        config
	store         store.Storage
	cacheStorage  cache.Storage
	logger        *zap.SugaredLogger
	mailer        mailer.Client
	authenticator auth.Authenticator
	rateLimiter   ratelimiter.Limiter
}

type config struct {
	addr        string
	db          dbConfig
	env         string
	apiURL      string
	mail        mailConfig
	frontendURL string
	auth        authConfig
	redisCfg    redisConfig
	rateLimiter ratelimiter.Config
}

type redisConfig struct {
	addr    string
	pw      string
	db      int
	enabled bool
}

type authConfig struct {
	basic basicConfig
	token tokenConfig
}

type tokenConfig struct {
	secret string
	exp    time.Duration
	iss    string
}

type dbConfig struct {
	addr         string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime  string
}
type basicConfig struct {
	user string
	pass string
}

type mailConfig struct {
	sendGrid  sendGridConfig
	fromEmail string
	exp       time.Duration
}

type sendGridConfig struct {
	apiKey string
}

func (app *application) mount() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{env.GetString("CORS_ALLOWED_ORIGIN", "http://localhost:5174")},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/posts", func(r chi.Router) {
		r.Post("/", app.createPostHandler)

		r.Route("/{postID}",
			func(r chi.Router) {
				r.Get("/", app.getPostHandler)

				// r.Patch("/", app.checkPostOwnership("moderator", app.updatePostHandler))
				// r.Delete("/", app.checkPostOwnership("admin", app.deletePostHandler))
			})
	})

	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", app.healthCheckHandler)

		r.Route("/{userID}", func(r chi.Router) {
			r.Get("/", app.getUserHandler)
		})

		r.Route("/users", func(r chi.Router) {
			r.Put("/activate/{token}", app.activateUserHandler)

			r.Route("/{userID}", func(r chi.Router) {
				r.Get("/", app.getUserHandler)
				r.Put("/follow", app.followUserHandler)
				r.Put("/unfollow", app.unfollowUserHandler)
			})

			r.Group(func(r chi.Router) {
				r.Get("/feed", app.getUserFeedHandler)
			})
		})
	})

	r.Route("/authentication", func(r chi.Router) {
		r.Post("/user", app.registerUserHandler)
		r.Post("/token", app.createTokenHandler)	})

	return r
}

func (app *application) run(mux http.Handler) error {
	// Create a new HTTP server with specified configurations
	srv := &http.Server{
		Addr:         app.config.addr,  // Set the address to listen on from the config
		Handler:      mux,              // Assign the HTTP handler (router/mux) to the server
		WriteTimeout: time.Second * 30, // Set the maximum duration for writing the response
		ReadTimeout:  time.Second * 10, // Set the maximum duration for reading the request
		IdleTimeout:  time.Minute,      // Set the maximum idle time for connections
	}

	// Channel to signal when the server should shut down
	shutdown := make(chan error)

	// Start a goroutine to listen for termination signals
	go func() {
		// Channel to receive OS signals
		quit := make(chan os.Signal, 1)

		// Notify the quit channel on SIGINT (Ctrl+C) or SIGTERM (termination request)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		// Wait for a signal to be received
		s := <-quit

		// Create a context with a 5-second timeout for the shutdown process
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel() // Ensure the context is cancelled after use

		// Log that a termination signal has been caught
		app.logger.Infow("signal caught", "signal", s.String())

		// Initiate the shutdown process of the server and send the result to the shutdown channel
		shutdown <- srv.Shutdown(ctx)
	}()

	// Log that the server has started, including address and environment info
	app.logger.Infow("server has started", "addr", app.config.addr, "env", app.config.env)

	// Start the server and listen for incoming requests
	err := srv.ListenAndServe()

	// Check if the error returned is due to the server being closed intentionally
	if !errors.Is(err, http.ErrServerClosed) {
		return err // If it's a different error, return it
	}

	// Wait for the shutdown signal to complete and retrieve any resulting error
	err = <-shutdown
	if err != nil {
		return err // Return any error that occurred during shutdown
	}

	// Log that the server has stopped
	app.logger.Infow("server has stopped", "addr", app.config.addr, "env", app.config.env)

	return nil // Indicate that the run method completed successfully
}
