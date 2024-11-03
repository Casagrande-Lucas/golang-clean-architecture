package migrations

import (
	"log"

	"github.com/Casagrande-Lucas/golang-clean-architecture/infrastructure/database"
	"github.com/Casagrande-Lucas/golang-clean-architecture/internal/domain"
)

func Migrate() {
	dbConn := database.GetDBInstance()

	if err := dbConn.AutoMigrate(&domain.User{}); err != nil {
		log.Fatal("failed to migrate database:", err)
		return
	}

	log.Println("Database migration completed.")
}
