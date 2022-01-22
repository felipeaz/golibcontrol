package jwt

import (
	"time"

	"github.com/FelipeAz/golibcontrol/infra/consumer/jwt/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/twinj/uuid"
)

// CreateToken generates a JWT Token
func CreateToken(userId uint, email, kid, secret string) (string, error) {
	td := model.TokenDetails{
		AtExpires:  time.Now().Add(time.Minute * 15).Unix(),
		AccessUuid: uuid.NewV4().String(),
	}

	// jwt.SigningMethodHS256 is the method used to generate the signature
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":         kid,
		"authorized":  true,
		"user_id":     userId,
		"email":       email,
		"access_uuid": td.AccessUuid,
		"exp":         td.AtExpires,
	})
	token.Header["kid"] = kid

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
