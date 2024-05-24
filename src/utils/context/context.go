package context

import (
	"ecommerce/src/dtos"
	"ecommerce/src/utils/db"

	"github.com/gin-gonic/gin"
)

type Context struct {
	*gin.Context
	*db.DB
	*dtos.User
}
