package main

import (
	"context"
	"net/http"

	"github.com/ishisora/go-practice/config"
)

func NewMux(ctx context.Context, cfg *config.Config) (*http.ServeMux, func(), error) {
	mux := http.NewServeMux()
	cleanup := func() {
	}
	return mux, cleanup, nil
}
