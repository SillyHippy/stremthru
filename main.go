package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/MunifTanjim/stremthru/internal/config"
	"github.com/MunifTanjim/stremthru/internal/db"
	"github.com/MunifTanjim/stremthru/internal/endpoint"
)

var mux *http.ServeMux

func init() {
	config.PrintConfig()

	mux = http.NewServeMux()

	endpoint.AddRootEndpoint(mux)
	endpoint.AddHealthEndpoints(mux)
	endpoint.AddProxyEndpoints(mux)
	endpoint.AddStoreEndpoints(mux)
	endpoint.AddStremioEndpoints(mux)

	database := db.Open()
	defer db.Close()
	db.Ping()
	RunSchemaMigration(database.URI)
}

// Handler is the exported serverless function entry point
func Handler(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	mux.ServeHTTP(w, r)
}
