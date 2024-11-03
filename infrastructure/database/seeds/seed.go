package seeds

import (
	"errors"
	"log"

	"github.com/Casagrande-Lucas/golang-clean-architecture/internal/domain"
	"github.com/Casagrande-Lucas/golang-clean-architecture/internal/repositories/mysql"
	"github.com/Casagrande-Lucas/golang-clean-architecture/internal/usecases/user_usecase"
	"github.com/google/uuid"
)

func Seed() {
	// Create an admin user object with predefined credentials
	adminUser := domain.User{
		ID:        uuid.New().String(),
		FirstName: "User",
		LastName:  "Admin",
		FullName:  "User Admin",
		Email:     "admin@admin.com",
		Password:  "test@123",
	}

	// Initialize the user use case with a MySQL repository implementation
	userUseCase := user_usecase.NewUserUseCase(mysql.NewUserRepository())

	// Attempt to register the admin user
	if err := userUseCase.Register(&adminUser); err != nil {
		if !errors.Is(err, user_usecase.ErrEmailAlreadyRegistered) {
			log.Fatalf("Failed to register admin user: %v", err)
			return
		}
		log.Println("Admin user already registered.")
	}

	log.Println("Seeded successfully.")
}
