package router

import (
	"github.com/FelipeAz/golibcontrol/build/server/account/router/build"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/account/users/handler"
	"github.com/FelipeAz/golibcontrol/internal/app/middleware"
	"github.com/gin-gonic/gin"
)

func Build(accountHandler handler.AccountHandler) error {
	router := gin.New()
	router.Use(middleware.CORSMiddleware())

	apiRg := router.Group("/api")
	vGroup := apiRg.Group("/v1")

	build.UserRoutes(vGroup, accountHandler)

	return router.Run(":8082")
}
