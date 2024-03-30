// package main

// import "net/http"

// type APIServer struct {
// 	addr string
// }

// func NewAPIServer(addr string) *APIServer {
// 	return &APIServer{addr: addr}
// }

// func (s *APIServer) Run() error {
// 	router := http.NewServeMux()
// 	router.HandleFunc("GET /users/{userID}", func(w http.ResponseWriter, r *http.Request) {
// 		userID := r.PathValue("userID")
// 		w.Write([]byte(userID))
// 	})

// 	server := http.Server{
// 		Addr:    s.addr,
// 		Handler: router,
// 	}
// 	return server.ListenAndServe()

// }
