package auth_usecase

import (
	"errors"
	"github.com/Casagrande-Lucas/golang-clean-architecture/internal/repositories/mysql"
	"github.com/Casagrande-Lucas/golang-clean-architecture/internal/usecases/jwt_usecase"
	"golang.org/x/crypto/bcrypt"
)

type IAuthUseCase interface {
	Login(email, password string) (string, error)
}

type authUseCase struct {
	userRepo   mysql.IUserRepository
	jwtUseCase jwt_usecase.IJWTUseCase
}

func NewAuthUseCase(userRepo mysql.IUserRepository, jwtUseCase jwt_usecase.IJWTUseCase) IAuthUseCase {
	return &authUseCase{
		userRepo:   userRepo,
		jwtUseCase: jwtUseCase,
	}
}

func (a *authUseCase) Login(email, password string) (string, error) {
	user, err := a.userRepo.FindByEmail(email)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("email or password is incorrect")
	}

	return a.jwtUseCase.GenerateToken(user.ID)
}
