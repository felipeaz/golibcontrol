package router

import (
	"github.com/FelipeAz/golibcontrol/build/router/account/router/routes"
	"github.com/FelipeAz/golibcontrol/infra/middleware"
	auth_handler "github.com/FelipeAz/golibcontrol/internal/app/account/auth/handler"
	account_handler "github.com/FelipeAz/golibcontrol/internal/app/account/users/handler"
	"github.com/gin-gonic/gin"
)

func Route(accountHandler account_handler.AccountHandler, authHandler auth_handler.AuthHandler) error {
	router := gin.New()
	router.Use(middleware.CORSMiddleware())

	apiRg := router.Group("/api")
	vGroup := apiRg.Group("/v1")

	routes.UserRoutes(vGroup, accountHandler)
	routes.AuthRoutes(vGroup, authHandler)

	return router.Run(":8082")
}
