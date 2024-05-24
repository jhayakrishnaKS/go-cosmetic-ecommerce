package handlers

import (
	"ecommerce/src/utils/db"
	"ecommerce/src/utils/context"

	"github.com/gin-gonic/gin"
)

func getContext(c *gin.Context) *context.Context{
	return &context.Context{
		Context:c, 
		DB: db.New(),
	}
}