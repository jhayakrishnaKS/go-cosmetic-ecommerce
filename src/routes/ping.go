package routes

import (
	"ecommerce/src/handlers"
	middlewares "ecommerce/src/middleware"

	"github.com/gin-gonic/gin"
)
 
func pingRoutes(router *gin.Engine) {
    router.Use(middlewares.CORSMiddleware())
    router.GET("/ping", handlers.Ping)
}
 