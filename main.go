package main

import (
	"context"
	"net/http"

	"github.com/MunifTanjim/stremthru/internal/config"
	"github.com/MunifTanjim/stremthru/internal/db"
	"github.com/MunifTanjim/stremthru/internal/endpoint"
)

var mux *http.ServeMux

func init() {
	// Initialize configuration
	config.PrintConfig()

	// Set up HTTP mux
	mux = http.NewServeMux()

	// Add application endpoints
	endpoint.AddRootEndpoint(mux)
	endpoint.AddHealthEndpoints(mux)
	endpoint.AddProxyEndpoints(mux)
	endpoint.AddStoreEndpoints(mux)
	endpoint.AddStremioEndpoints(mux)

	// Initialize the database
	database := db.Open()
	defer db.Close()
	db.Ping()
	RunSchemaMigration(database.URI)
}

// Handler is the exported serverless function for Vercel
func Handler(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	mux.ServeHTTP(w, r)
}
