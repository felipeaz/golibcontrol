package jwt

import (
	"net/http"
	"os"
	"time"

	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/constants/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/twinj/uuid"
)

// CreateToken generates a JWT Token
func CreateToken(userId uint) (model.TokenDetails, *errors.ApiError) {
	td := model.TokenDetails{
		AtExpires:   time.Now().Add(time.Minute * 15).Unix(),
		AccessUuid:  uuid.NewV4().String(),
		RtExpires:   time.Now().Add(time.Hour * 24 * 7).Unix(),
		RefreshUuid: uuid.NewV4().String(),
	}

	// Payload
	atClaims := jwt.MapClaims{
		"authorized": true,
		"user_id":    userId,
		"exp":        td.AtExpires,
	}

	// Secret key used on signature crypto
	jwtSecret := []byte(os.Getenv("JWT_SECRET"))

	// jwt.SigningMethodHS256 is the method used to generate the signature
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	aToken, err := at.SignedString(jwtSecret)
	if err != nil {
		return model.TokenDetails{}, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: errors.JWTTokenCreationFailMessage,
			Error:   err.Error(),
		}
	}

	// Create the refresh token
	rtClaims := jwt.MapClaims{
		"authorized": true,
		"user_id":    userId,
		"exp":        td.RtExpires,
	}

	jwtRefreshSecret := []byte(os.Getenv("JWT_REFRESH_SECRET"))
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	rToken, err := rt.SignedString(jwtRefreshSecret)
	if err != nil {
		return model.TokenDetails{}, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: errors.JWTTokenCreationFailMessage,
			Error:   err.Error(),
		}
	}

	td.AccessToken = aToken
	td.RefreshToken = rToken

	return td, nil
}
