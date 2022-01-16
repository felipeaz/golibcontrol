package middleware

import (
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"access-control-allow-origin", "access-control-allow-headers"},
	})
}
