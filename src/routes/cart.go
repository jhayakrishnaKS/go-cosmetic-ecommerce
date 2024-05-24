package routes

import (
	"ecommerce/src/handlers"
	middlewares "ecommerce/src/middleware"

	"github.com/gin-gonic/gin"
)

func CartRoutes(router *gin.Engine){
	router.GET("/cart",middlewares.WithAuth(handlers.GetAllCart))
	router.POST("/cart",middlewares.WithAuth(handlers.AddToCart))
	router.DELETE("/cart/:id",middlewares.WithAuth(handlers.DeleteCart))
}