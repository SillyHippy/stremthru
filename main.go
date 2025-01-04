package main

import (
	"log"
	"net/http"
	"github.com/MunifTanjim/stremthru/internal/config"
	"github.com/MunifTanjim/stremthru/internal/db"
	"github.com/MunifTanjim/stremthru/internal/endpoint"
)

var (
	mux      *http.ServeMux
	database *db.Database
)

func init() {
	// Initialize configuration
	config.PrintConfig()

	// Create ServeMux and add endpoints
	mux = http.NewServeMux()
	endpoint.AddRootEndpoint(mux)
	endpoint.AddHealthEndpoints(mux)
	endpoint.AddProxyEndpoints(mux)
	endpoint.AddStoreEndpoints(mux)
	endpoint.AddStremioEndpoints(mux)

	// Connect to the database
	var err error
	database, err = db.Open()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Test database connection
	if err := db.Ping(); err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	// Run schema migrations
	if err := Run
