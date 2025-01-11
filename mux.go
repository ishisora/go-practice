package main

import (
	"context"
	"net/http"

	"github.com/ishisora/go-practice/config"
)

func NewMux(ctx context.Context, cfg *config.Config) (*http.ServeMux, func(), error) {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("HealthCheck: OK"))
	})

	cleanup := func() {
	}
	return mux, cleanup, nil
}
