package router

import (
	"github.com/FelipeAz/golibcontrol/build/router/account/router/routes"
	"github.com/FelipeAz/golibcontrol/infra/middleware"
	authHandler "github.com/FelipeAz/golibcontrol/internal/app/account/auth/handler"
	accountHandler "github.com/FelipeAz/golibcontrol/internal/app/account/users/handler"
	"github.com/gin-gonic/gin"
)

type Handlers struct {
	AccountHandler accountHandler.AccountHandler
	AuthHandler    authHandler.AuthHandler
}

func Route(handlers Handlers) error {
	router := gin.New()
	router.Use(middleware.CORSMiddleware())

	apiRg := router.Group("/api")
	vGroup := apiRg.Group("/v1")

	routes.UserRoutes(vGroup, handlers.AccountHandler)
	routes.AuthRoutes(vGroup, handlers.AuthHandler)

	return router.Run(":8082")
}
