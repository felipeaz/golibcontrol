package middleware

import (
	"net/http"

	"github.com/FelipeAz/golibcontrol/platform/jwt"
	"github.com/gin-gonic/gin"
)

// TokenMiddleware contains JWTAuth
type TokenMiddleware struct {
	JWTAuth *jwt.Auth
}

// NewTokenMiddleware returns an instance of TokenMiddleware
func NewTokenMiddleware(jwtAuth *jwt.Auth) *TokenMiddleware {
	return &TokenMiddleware{
		JWTAuth: jwtAuth,
	}
}

// TokenAuth authenticates the bearer token from request
func (tm TokenMiddleware) TokenAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := tm.JWTAuth.TokenValid(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}

		c.Next()
	}
}
