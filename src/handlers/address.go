package handlers

import (
	"ecommerce/src/dtos"
	middlewares "ecommerce/src/middleware"
	"ecommerce/src/services/address"
	"ecommerce/src/utils/context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllAddress(c *context.Context) {
	userID := c.User.ID
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User ID is required",
		})
		return
	}

	addresses, err := address.NewAddress().GetAllAddress(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"addresses": addresses,
	})
}

func CreateAddress(c *context.Context) {
	req := &dtos.AddressReq{}
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

	err := address.NewAddress().AddAddress(c, req, token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "address added successfully",
	})
}

func UpdateAddress(c *context.Context) {
	req := &dtos.UpdateAddressReq{}
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

	err := address.NewAddress().UpdateAddress(c, req, token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Address updated successfully",
	})
}

func DeleteAddress(c *context.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID parameter is required",
		})
		return
	}

	err := address.NewAddress().DeleteAddress(c, id)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "address deleted successfully",
	})
}
