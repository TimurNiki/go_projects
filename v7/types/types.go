package main

import (
	"net/http"
	"sync")

type Server struct {
	subscriberMessageBuffer int
	mux                     http.ServeMux
	subscribersMu           sync.Mutex
	subscribers             map[*subscriber]struct{}
}

type Subscriber struct {
	msgs chan []byte
}