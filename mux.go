package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/ishisora/go-practice/config"
)

func NewMux(ctx context.Context, cfg *config.Config) (http.Handler, func(), error) {
	r := chi.NewRouter()
	r.Get("/health", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintln(w, "HealthCheck: OK")
	})

	cleanup := func() {
	}

	return r, cleanup, nil
}
