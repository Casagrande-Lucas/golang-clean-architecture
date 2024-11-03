package main

import (
	"log"

	"github.com/Casagrande-Lucas/golang-clean-architecture/infrastructure/database/migrations"
	"github.com/Casagrande-Lucas/golang-clean-architecture/infrastructure/database/seeds"
	"github.com/Casagrande-Lucas/golang-clean-architecture/internal/delivery/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file, if it exists
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on environment variables")
	}

	// Run database migrations to update schema
	migrations.Migrate()

	// Seed the database with initial data (e.g., admin user)
	seeds.Seed()

	// Initialize the Gin router and register all application routes
	r := gin.Default()
	routes.RegisterRoutes(r)

	// Start the HTTP server on port 8080
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
