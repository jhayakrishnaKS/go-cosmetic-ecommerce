package handlers

import (
	"ecommerce/src/dtos"
	"ecommerce/src/services/orderstatus"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateStatus(c *gin.Context) {
	var req dtos.OrderStatusReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := orderstatus.NewStatus().CreateStatus(getContext(c), &req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Order status created successfully",
	})
}
