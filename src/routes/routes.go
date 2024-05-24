package routes

import (
	middlewares "ecommerce/src/middleware"
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	router := gin.Default()

	router.Use(middlewares.CORSMiddleware())

	pingRoutes(router)
	ProductRoutes(router)
	RoleRoutes(router)
	UserRoutes(router)
	OrderRoutes(router)
	CartRoutes(router)
	AddressRoutes(router)
	

	return router
}
