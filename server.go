package main

import (
	"context"
	"net"
	"net/http"
)

type Server struct {
	srv *http.Server
	l   net.Listener
}

func NewServer(l net.Listener, mux *http.ServeMux) *Server {
	return &Server{
		srv: &http.Server{Handler: mux},
		l:   l,
	}
}

func (s *Server) Run(ctx context.Context) error {
	server := &http.Server{
		Handler: s.srv.Handler,
	}
	go func() {
		<-ctx.Done()
		server.Close()
	}()
	return server.Serve(s.l)
}
