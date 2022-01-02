package router

import (
	"github.com/FelipeAz/golibcontrol/build/server/account/router/build"
	auth_handler "github.com/FelipeAz/golibcontrol/internal/app/account/auth/handler"
	account_handler "github.com/FelipeAz/golibcontrol/internal/app/account/users/handler"
	"github.com/FelipeAz/golibcontrol/internal/app/middleware"
	"github.com/gin-gonic/gin"
)

func Build(accountHandler account_handler.AccountHandler, authHandler auth_handler.AuthHandler) error {
	router := gin.New()
	router.Use(middleware.CORSMiddleware())

	apiRg := router.Group("/api")
	vGroup := apiRg.Group("/v1")

	build.UserRoutes(vGroup, accountHandler)
	build.AuthRoutes(vGroup, authHandler)

	return router.Run(":8082")
}
