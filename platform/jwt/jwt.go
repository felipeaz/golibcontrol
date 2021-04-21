package jwt

import (
	"net/http"
	"os"
	"time"

	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/dgrijalva/jwt-go"
)

// CreateToken generates a JWT Token
func CreateToken(userId uint) (string, *errors.ApiError) {
	atClaims := jwt.MapClaims{
		"authorized": true,
		"user_id":    userId,
		"exp":        time.Now().Add(time.Minute * 15).Unix(),
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: errors.JWTTokenCreationFailMessage,
			Error:   err.Error(),
		}
	}
	return token, nil
}
