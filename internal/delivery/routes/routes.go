package routes

import (
	"github.com/Casagrande-Lucas/golang-clean-architecture/internal/delivery/http"
	"github.com/Casagrande-Lucas/golang-clean-architecture/internal/middlewares"
	"github.com/Casagrande-Lucas/golang-clean-architecture/internal/repositories/mysql"
	"github.com/Casagrande-Lucas/golang-clean-architecture/internal/usecases/auth_usecase"
	"github.com/Casagrande-Lucas/golang-clean-architecture/internal/usecases/jwt_usecase"
	"github.com/Casagrande-Lucas/golang-clean-architecture/internal/usecases/user_usecase"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	jwtMiddleware := middlewares.NewJWTMiddleware()

	authHandler := http.NewAuthHandler(auth_usecase.NewAuthUseCase(mysql.NewUserRepository(), jwt_usecase.NewJWTUseCase()))
	authHandler.RegisterRoutes(r)

	userHandler := http.NewUserHandler(user_usecase.NewUserUseCase(mysql.NewUserRepository()), jwtMiddleware)
	userHandler.RegisterRoutes(r)
}
