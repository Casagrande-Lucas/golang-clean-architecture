package http

import (
	"errors"
	"github.com/Casagrande-Lucas/golang-clean-architecture/internal/domain"
	"github.com/Casagrande-Lucas/golang-clean-architecture/internal/helpers"
	"github.com/Casagrande-Lucas/golang-clean-architecture/internal/middlewares"
	"github.com/Casagrande-Lucas/golang-clean-architecture/internal/types"
	"github.com/Casagrande-Lucas/golang-clean-architecture/internal/usecases/user_usecase"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type IUserHandler interface {
	RegisterRoutes(r *gin.Engine)
	Register(c *gin.Context)
	GetUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	ResetPassword(c *gin.Context)
}

type userHandler struct {
	userUseCase   user_usecase.IUserUseCase
	jwtMiddleware middlewares.IJWTMiddleware
	validator     *validator.Validate
}

func NewUserHandler(us user_usecase.IUserUseCase, middleware middlewares.IJWTMiddleware) IUserHandler {
	return &userHandler{
		userUseCase:   us,
		jwtMiddleware: middleware,
		validator:     validator.New(),
	}
}

func (uh *userHandler) RegisterRoutes(r *gin.Engine) {
	userGroup := r.Group("/api/v1/users", uh.jwtMiddleware.Middleware())
	{
		userGroup.POST("/", uh.Register)
		userGroup.GET("/", uh.GetAllUsers)
		userGroup.GET("/:id", uh.GetUser)
		userGroup.PUT("/:id", uh.UpdateUser)
		userGroup.DELETE("/:id", uh.DeleteUser)
		userGroup.POST("/reset_password", uh.ResetPassword)
	}
}

func (uh *userHandler) Register(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := uh.validator.Struct(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := uh.userUseCase.Register(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "user": user})
}

func (uh *userHandler) GetAllUsers(c *gin.Context) {
	offset, _ := strconv.Atoi(c.Query("offset"))
	limit, _ := strconv.Atoi(c.Query("limit"))

	if offset < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "offset cannot be negative"})
		return
	}

	if limit <= 0 || limit > 100 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "limit cannot be negative, zero or greater than 100"})
		return
	}

	users, total, err := uh.userUseCase.GetAllUsers(offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	currentPage := (offset / limit) + 1

	totalPages := int((total + int64(limit) - 1) / int64(limit))

	response := types.UserPaginationResponse{
		Data:        users,
		Total:       total,
		PageSize:    limit,
		CurrentPage: currentPage,
		TotalPages:  totalPages,
	}

	c.JSON(http.StatusOK, response)
}

func (uh *userHandler) GetUser(c *gin.Context) {
	id := c.Param("id")
	if err := helpers.IsValidUUIDv4(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := uh.userUseCase.GetUser(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Get user", "user": user})
}

func (uh *userHandler) UpdateUser(c *gin.Context) {
	userID := c.Param("id")
	if err := helpers.IsValidUUIDv4(userID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user domain.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := uh.validator.Struct(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.ID = userID
	if err := uh.userUseCase.UpdateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User updated successfully"})
}

func (uh *userHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if err := helpers.IsValidUUIDv4(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := uh.userUseCase.DeleteUser(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Delete user"})
}

func (uh *userHandler) ResetPassword(c *gin.Context) {
	userID, err := helpers.GetUserIDInContextRequest(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	if err := helpers.IsValidUUIDv4(userID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var resetPasswordRequest struct {
		OldPassword        string `json:"old_password" validate:"required"`
		NewPassword        string `json:"new_password" validate:"required"`
		ConfirmNewPassword string `json:"confirm_new_password" validate:"required"`
	}

	if err := c.ShouldBind(&resetPasswordRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := uh.validator.Struct(&resetPasswordRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := uh.userUseCase.ResetPassword(userID, resetPasswordRequest.OldPassword, resetPasswordRequest.NewPassword); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password reset"})
}
