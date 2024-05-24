package handlers

import (
	"ecommerce/src/dtos"
	middlewares "ecommerce/src/middleware"
	"ecommerce/src/services/cart"
	"ecommerce/src/utils/context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllCart(c *context.Context){
	userID := c.User.ID
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User ID is required",
		})
		return
	}

	carts,err:=cart.NewCart().GetAllItems(c,userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"carts": carts,
	})
}


func AddToCart(c *context.Context) {
	req := &dtos.CartReq{}
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

	err := cart.NewCart().AddToCart(c, req, token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Cart added successfully",
	})
}

func DeleteCart(c *context.Context){
	id:=c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID parameter is required",
		})
		return
	}

	err:=cart.NewCart().Delete(c,id)
	if err!=nil{
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "cart deleted successfully",
	})
}
