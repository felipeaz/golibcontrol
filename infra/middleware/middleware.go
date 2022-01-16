package middleware

import (
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	"strings"
)

type Middleware struct {
	allowedOrigins []string
	allowedHeaders []string
}

func New(allowedOrigins, allowedHeaders string) *Middleware {
	return &Middleware{
		allowedOrigins: strings.Split(strings.ReplaceAll(allowedOrigins, " ", ""), ","),
		allowedHeaders: strings.Split(strings.ReplaceAll(allowedHeaders, " ", ""), ","),
	}
}

func (m Middleware) CORSMiddleware() gin.HandlerFunc {
	return cors.New(cors.Options{
		AllowedOrigins: m.allowedOrigins,
		AllowedHeaders: m.allowedHeaders,
	})
}
