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
	// Initialize your app's configuration
	config.PrintConfig()

	// Create a new ServeMux and add endpoints
	mux = http.NewServeMux()
	endpoint.AddRootEndpoint(mux)
	endpoint.AddHealthEndpoints(mux)
	endpoint.AddProxyEndpoints(mux)
	endpoint.AddStoreEndpoints(mux)
	endpoint.AddStremioEndpoints(mux)

	// Connect to the database
	database := db.Open()
	defer db.Close()
	db.Ping()

	// Run any necessary schema migrations
	RunSchemaMigration(database.URI)
}

// Handler is the exported function required by Vercel
func Handler(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	mux.ServeHTTP(w, r) // Pass requests to your ServeMux
}
