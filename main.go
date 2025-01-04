package main

import (
	"context"
	"net/http"
	"github.com/MunifTanjim/stremthru/internal/config"
	"github.com/MunifTanjim/stremthru/internal/db"
	"github.com/MunifTanjim/stremthru/internal/endpoint"
)

var mux *http.ServeMux

func setupMux() *http.ServeMux {
	mux := http.NewServeMux()
	endpoint.AddRootEndpoint(mux)
	endpoint.AddHealthEndpoints(mux)
	endpoint.AddProxyEndpoints(mux)
	endpoint.AddStoreEndpoints(mux)
	endpoint.AddStremioEndpoints(mux)
	return mux
}

func setupDatabase() {
	database := db.Open()
	defer db.Close() // Ensure proper cleanup when the application terminates
	db.Ping()
	RunSchemaMigration(database.URI)
}

func init() {
	// Initialize configuration
	config.PrintConfig()

	// Setup database and routes
	setupDatabase()
	mux = setupMux()
}

// Handler is the exported function required by Vercel
func Handler(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	mux.ServeHTTP(w, r)
}
