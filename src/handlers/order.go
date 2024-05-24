package handlers

import (
	"ecommerce/src/dtos"
	middlewares "ecommerce/src/middleware"
	"ecommerce/src/services/order"
	"ecommerce/src/utils/context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllOrder(c *context.Context){
	userID := c.User.ID
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User ID is required",
		})
		return
	}

	orders,err:=order.NewOrder().GetAllOrder(c,userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"orders": orders,
	})
}

func Checkout(c *context.Context){
	req:=&dtos.OrderReq{}
	if err := c.Bind(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	token := middlewares.GetBearerToken(c.Context)
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "No bearer token provided",
		})
		return
	}

	err:=order.NewOrder().Checkout(c,req,token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "order added successfully",
	})
}

func Update(c *context.Context) {
	if c.User.RoleID != "947019b0-0970-4dd2-9dfa-8320cd871e06" {
        c.JSON(http.StatusForbidden, gin.H{
            "error": "You do not have permission to update order status",
        })
        return
    }
	orderId := c.Param("id")
	if orderId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "order ID is required",
		})
		return
	}

	var updateReq dtos.OrderReq
	if err := c.Bind(&updateReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	orderService := order.NewOrder()

	err := orderService.UpdateStatus(c, orderId, &updateReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Order status updated successfully",
	})
}