package main

import (
	"log"

	"github.com/Kevinmajesta/backend_library/configs"
	"github.com/Kevinmajesta/backend_library/internal/builder"
	"github.com/Kevinmajesta/backend_library/pkg/postgres"
	"github.com/Kevinmajesta/backend_library/pkg/server"
)

func main() {
	// Load environment variables
	cfg, err := configs.NewConfig(".env")
	checkError(err)

	// Init PostgreSQL DB
	db, err := postgres.InitPostgres(&cfg.Postgres)
	checkError(err)

	// Build Echo route groups
	publicRoutes := builder.BuildPublicRoutes(db)
	privateRoutes := builder.BuildPrivateRoutes(db)

	// Start server
	srv := server.NewServer("app", publicRoutes, privateRoutes)
	srv.Run()
}

func checkError(err error) {
	if err != nil {
		log.Fatalf("Fatal error: %v", err)
	}
}
