package helpers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func IsValidUUIDv4(u string) error {
	id, err := uuid.Parse(u)
	if err != nil {
		return err
	}

	if ok := id.Version() == uuid.Version(4); !ok {
		return errors.New("invalid UUIDv4")
	}

	return nil
}

func GetUserIDInContextRequest(c *gin.Context) (string, error) {
	userID, exists := c.Get("userID")
	if !exists {
		return "", errors.New("unauthorized")
	}

	return userID.(string), nil
}
