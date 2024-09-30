package main

import "net/http"

type userKey string

const userCtx userKey = "user" 

func getUserFromContext(r *http.Request) *store.User {
	user, _ := r.Context().Value(userCtx).(*store.User)
	return user
}