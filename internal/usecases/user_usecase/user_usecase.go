package user_usecase

import (
	"errors"
	"github.com/Casagrande-Lucas/golang-clean-architecture/internal/domain"
	"github.com/Casagrande-Lucas/golang-clean-architecture/internal/helpers"
	"github.com/Casagrande-Lucas/golang-clean-architecture/internal/repositories/mysql"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type IUserUseCase interface {
	Register(user *domain.User) error
	GetUser(id string) (*domain.User, error)
	GetAllUsers(offset, limit int) (*[]domain.User, int64, error)
	UpdateUser(user *domain.User) error
	DeleteUser(id string) error
	ResetPassword(userID, oldPassword, newPassword string) error
}

var ErrEmailAlreadyRegistered = errors.New("email already registered")

type userUseCase struct {
	userRepo mysql.IUserRepository
}

func NewUserUseCase(userRepo mysql.IUserRepository) IUserUseCase {
	return &userUseCase{userRepo: userRepo}
}

func (uc *userUseCase) Register(user *domain.User) error {
	_, err := uc.userRepo.FindByEmail(user.Email)
	if err == nil {
		return ErrEmailAlreadyRegistered
	}

	user.ID = uuid.NewString()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	return uc.userRepo.Create(user)
}

func (uc *userUseCase) GetUser(id string) (*domain.User, error) {
	if err := helpers.IsValidUUIDv4(id); err != nil {
		return nil, err
	}
	return uc.userRepo.FindByID(id)
}

func (uc *userUseCase) GetAllUsers(offset, limit int) (*[]domain.User, int64, error) {
	return uc.userRepo.FindAll(offset, limit)
}

func (uc *userUseCase) UpdateUser(user *domain.User) error {
	userFind, err := uc.userRepo.FindByID(user.ID)
	if err != nil {
		return err
	}

	user.CreatedAt = userFind.CreatedAt
	user.Password = userFind.Password
	user.Active = true

	if user.FullName == "" {
		user.FullName = user.FirstName + " " + user.LastName
	}

	return uc.userRepo.Update(user)
}

func (uc *userUseCase) DeleteUser(id string) error {
	return uc.userRepo.Delete(id)
}

func (uc *userUseCase) ResetPassword(userID, oldPassword, newPassword string) error {
	user, err := uc.userRepo.FindByID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		}
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldPassword)); err != nil {
		return errors.New("invalid old password")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	if err := uc.userRepo.Update(user); err != nil {
		return err
	}

	return nil
}
