package routes

import (
	"ecommerce/src/handlers"
	middlewares "ecommerce/src/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	router.POST("/register", handlers.Register)
	router.GET("/user", handlers.GetAccessFromRefreshToken)
	router.POST("/login", handlers.Login)
	router.GET("/users", middlewares.WithAuth(handlers.GetUsers))
	router.GET("/account", middlewares.WithAuth(handlers.GetAccountByUsingAccessToken))
}
