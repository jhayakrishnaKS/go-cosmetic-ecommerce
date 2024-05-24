package handlers

import (
	"net/http"
	"ecommerce/src/dtos"
	"ecommerce/src/services/roles"

	"github.com/gin-gonic/gin"
)

func RegisterRoles(c *gin.Context){
	req:=&dtos.RoleReq{}
	if err:=c.Bind(req);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":err.Error(),
		})
		return
	}

	err:=roles.New().RegisterRoles(getContext(c),req)
	if err!=nil{
		c.JSON(http.StatusUnauthorized,gin.H{
			"error":err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Role created successfully",
	})
}