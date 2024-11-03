package mysql

import (
	"github.com/Casagrande-Lucas/golang-clean-architecture/infrastructure/database"
	"github.com/Casagrande-Lucas/golang-clean-architecture/internal/domain"
	"gorm.io/gorm"
)

type IUserRepository interface {
	Create(user *domain.User) error
	FindByID(id string) (*domain.User, error)
	FindByEmail(email string) (*domain.User, error)
	FindAll(offset, limit int) (*[]domain.User, int64, error)
	Update(user *domain.User) error
	Delete(id string) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository() IUserRepository {
	return &userRepository{db: database.GetDBInstance()}
}

func (r *userRepository) Create(user *domain.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) FindByID(id string) (*domain.User, error) {
	var user domain.User
	if err := r.db.Where("id = ? AND active = true", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindByEmail(email string) (*domain.User, error) {
	var user domain.User
	if err := r.db.Where("email = ? AND active = true", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindAll(offset, limit int) (*[]domain.User, int64, error) {
	var users []domain.User
	var total int64

	if err := r.db.Model(&domain.User{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := r.db.
		Order("created_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&users, "active = true").Error
	if err != nil {
		return nil, 0, err
	}

	return &users, total, nil
}

func (r *userRepository) Update(user *domain.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) Delete(id string) error {
	return r.db.Delete(&domain.User{}, "id = ?", id).Error
}
