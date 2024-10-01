package main

import (
	"net/http"
	"v5/internal/store"
)

type postKey string

const postCtx postKey = "post"

type CreatePostPayload struct {
	Title   string   `json:"title" validate:"required,max=100"`
	Content string   `json:"content" validate:"required,max=1000"`
	Tags    []string `json:"tags"`
}

func (app *application) createPostHandler(w http.ResponseWriter, r *http.Request) {
	var payload CreatePostPayload
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err:=Validate.Struct(payload);err!=nil{
		app.badRequestResponse(w, r, err)
		return
	}

	user:=getUserFromContext(r)

	post:=&store.Post{
		Title: payload.Title,
		Content:payload.Content,
		Tags: payload.Tags,
		UserID: user.ID,
	}
	
	ctx := r.Context()

	if err:=app.store.Posts.Create(ctx, post); err!=nil{
		app.internalServerError(w,r, err)
		return
	}
	if err := app.jsonResponse(w, http.StatusCreated, post); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}
