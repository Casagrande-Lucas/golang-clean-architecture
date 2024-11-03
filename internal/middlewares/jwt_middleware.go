package middlewares

import (
	"github.com/Casagrande-Lucas/golang-clean-architecture/internal/usecases/jwt_usecase"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type IJWTMiddleware interface {
	Middleware() gin.HandlerFunc
}

type jwtMiddleware struct {
	jwtUseCase jwt_usecase.IJWTUseCase
}

func NewJWTMiddleware() IJWTMiddleware {
	return &jwtMiddleware{jwtUseCase: jwt_usecase.NewJWTUseCase()}
}

func (m *jwtMiddleware) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			return
		}

		tokenString := strings.TrimSpace(strings.Replace(authHeader, "Bearer", "", 1))
		token, err := m.jwtUseCase.ValidateToken(tokenString)
		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			return
		}

		c.Set("userID", claims["user_id"])

		c.Next()
	}
}
