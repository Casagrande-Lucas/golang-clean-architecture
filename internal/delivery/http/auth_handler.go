package http

import (
	"github.com/Casagrande-Lucas/golang-clean-architecture/internal/usecases/auth_usecase"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type IAuthHandler interface {
	RegisterRoutes(r *gin.Engine)
	Login(c *gin.Context)
}

type authHandler struct {
	authUseCase auth_usecase.IAuthUseCase
	validator   *validator.Validate
}

func NewAuthHandler(authUseCase auth_usecase.IAuthUseCase) IAuthHandler {
	return &authHandler{
		authUseCase: authUseCase,
		validator:   validator.New(),
	}
}

func (h *authHandler) RegisterRoutes(r *gin.Engine) {
	userGroup := r.Group("/api/v1/auth")
	{
		userGroup.POST("/login", h.Login)
	}
}

func (h *authHandler) Login(c *gin.Context) {
	var credentials struct {
		Email    string `json:"email" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	if err := c.ShouldBind(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "token": nil})
		return
	}

	if err := h.validator.Struct(credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "token": nil})
		return
	}

	token, err := h.authUseCase.Login(credentials.Email, credentials.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error(), "token": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})
}
