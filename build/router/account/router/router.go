package router

import (
	"github.com/FelipeAz/golibcontrol/build/router/account/router/routes"
	"github.com/FelipeAz/golibcontrol/infra/middleware"
	authHandler "github.com/FelipeAz/golibcontrol/internal/app/account/auth/handler"
	accountHandler "github.com/FelipeAz/golibcontrol/internal/app/account/users/handler"
	"github.com/gin-gonic/gin"
)

func Route(
	accountHandler accountHandler.AccountHandler,
	authHandler authHandler.AuthHandler,
	mwr *middleware.Middleware) error {
	router := gin.New()
	router.Use(mwr.CORSMiddleware())

	apiRg := router.Group("/api")
	vGroup := apiRg.Group("/v1")

	routes.UserRoutes(vGroup, accountHandler)
	routes.AuthRoutes(vGroup, authHandler)

	return router.Run(":8082")
}
