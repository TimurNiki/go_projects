package main

import "net/http"

type httpServer struct {
	addr string
}

func NewHttpServer(addr string) *httpServer {
	return &httpServer{
		addr: addr,
	}
}

func (s *httpServer) Run() error {
	router := http.NewServeMux()

	orderHandler.RegisterRouter(router)
	return http.ListenAndServe(s.addr, router)
}