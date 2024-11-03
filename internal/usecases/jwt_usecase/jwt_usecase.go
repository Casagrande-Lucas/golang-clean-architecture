package jwt_usecase

import (
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

type IJWTUseCase interface {
	GenerateToken(userID string) (string, error)
	ValidateToken(tokenString string) (*jwt.Token, error)
}

type jwtUseCase struct {
	secretKey string
	issuer    string
}

type jwtCustomClaim struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

func NewJWTUseCase() IJWTUseCase {
	return &jwtUseCase{
		secretKey: os.Getenv("JWT_SECRET_KEY"),
		issuer:    os.Getenv("JWT_ISSUER"),
	}
}

func (j *jwtUseCase) GenerateToken(userID string) (string, error) {
	claims := &jwtCustomClaim{
		userID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			Issuer:    j.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(j.secretKey))
}

func (j *jwtUseCase) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(j.secretKey), nil
	})
	return token, err
}
