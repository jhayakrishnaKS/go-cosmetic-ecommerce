package routes

import (
	"ecommerce/src/handlers"
	middlewares "ecommerce/src/middleware"

	"github.com/gin-gonic/gin"
)

func AddressRoutes(router *gin.Engine) {
	router.GET("/address",middlewares.WithAuth(handlers.GetAllAddress))
	router.POST("/address", middlewares.WithAuth(handlers.CreateAddress))
	router.PUT("/address/:id",middlewares.WithAuth(handlers.UpdateAddress))
	router.DELETE("address/:id",middlewares.WithAuth(handlers.DeleteAddress))
}
