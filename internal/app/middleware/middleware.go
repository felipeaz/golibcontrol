package middleware

import (
	"net/http"

	"github.com/FelipeAz/golibcontrol/infra/jwt"
	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/gin-gonic/gin"
)

const (
	AuthenticationRequiredMessage = "Authentication Required"
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

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// TokenAuth authenticates the bearer token from request
func (tm TokenMiddleware) TokenAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := tm.JWTAuth.TokenValid(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, errors.ApiError{
				Message: AuthenticationRequiredMessage,
				Error:   err.Error(),
			})
			c.Abort()
			return
		}
		usrErr := tm.JWTAuth.UserLoggedIn(c.Request)
		if usrErr != nil {
			c.JSON(http.StatusUnauthorized, errors.ApiError{
				Message: AuthenticationRequiredMessage,
				Error:   usrErr.Error(),
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
