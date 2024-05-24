package routes

import (
	"ecommerce/src/handlers"
	middlewares "ecommerce/src/middleware"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(router *gin.Engine) {
	router.GET("/orders", middlewares.WithAuth(handlers.GetAllOrder))
	router.POST("/order", middlewares.WithAuth(handlers.Checkout))
	router.POST("/create/orderstatus", handlers.CreateStatus)
	router.PUT("/admin/orderstatus/:id", middlewares.WithAuth(handlers.Update))
}
