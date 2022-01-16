package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"
	"time"
)

func CORSMiddleware() gin.HandlerFunc {
	return cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	})
}
