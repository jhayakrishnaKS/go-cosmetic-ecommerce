package routes

import (
	"ecommerce/src/handlers"
	middlewares "ecommerce/src/middleware"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(router *gin.Engine) {
	
	router.POST("/admin/products", middlewares.WithAuth(handlers.Products))
	router.GET("/admin/products",middlewares.WithAuth(handlers.GetAllProducts))
	router.GET("/products",middlewares.WithAuth(handlers.GetAllProducts))
	router.PUT("/admin/products/:id",middlewares.WithAuth(handlers.UpdateProducts))
	router.DELETE("/admin/products/:id",middlewares.WithAuth(handlers.DeleteProducts))
}
