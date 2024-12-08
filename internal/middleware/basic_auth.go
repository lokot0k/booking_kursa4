package middleware

import (
	"errors"
	"meeting-room-booking/internal/domain"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type UserRepository interface {
	GetAll() ([]domain.User, error)
	GetByID(id int) (*domain.User, error)
	GetByUsername(username string) (*domain.User, error)
	Create(user domain.User) error
	Update(user domain.User) error
	Delete(id int) error
}

func BasicAuthMiddleware(userRepo UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Получаем заголовок Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		// Проверяем, что заголовок начинается с "Basic"
		if !strings.HasPrefix(authHeader, "Basic ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header"})
			c.Abort()
			return
		}

		// Декодируем базовую аутентификацию
		payload, err := decodeBasicAuth(authHeader)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header"})
			c.Abort()
			return
		}

		// Извлекаем username и password
		username := payload[0]
		password := payload[1]

		// Проверяем пользователя в базе данных
		user, err := userRepo.GetByUsername(username)
		if err != nil || user.Password != password {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
			c.Abort()
			return
		}

		// Устанавливаем контекст пользователя
		c.Set("user", user)
		c.Next()
	}
}

func decodeBasicAuth(authHeader string) ([]string, error) {
	// Извлекаем закодированную часть
	credentials := strings.TrimPrefix(authHeader, "Basic ")

	// Разделяем на username и password
	parts := strings.SplitN(credentials, ":", 2)
	if len(parts) != 2 {
		return nil, errors.New("invalid basic auth format")
	}

	return parts, nil
}
