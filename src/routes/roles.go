package routes

import (
	"ecommerce/src/handlers"
	middlewares "ecommerce/src/middleware"

	"github.com/gin-gonic/gin"
)

func RoleRoutes(router *gin.Engine){
	router.Use(middlewares.CORSMiddleware())
	router.POST("/roles",handlers.RegisterRoles)
}