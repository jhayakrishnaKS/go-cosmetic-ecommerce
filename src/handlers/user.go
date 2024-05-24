package handlers

import (
	"ecommerce/src/dtos"
	"ecommerce/src/services/user"
	"ecommerce/src/utils/context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context){
	req:=&dtos.RegisterReq{}
	if err:=c.Bind(req);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":err.Error(),
		})
		return
	}
	err:=user.NewUsers().Register(getContext(c),req)
	if err!=nil{
		c.JSON(http.StatusUnauthorized,gin.H{
			"error":err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Registeration successfull",
	})
}

func GetUsers(c *context.Context) {
	users, err := user.NewUsers().GetUsers(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func Login(c *gin.Context) {
	req := &dtos.LoginReq{}
	if err := c.Bind(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userService := user.NewUsers()
	loginRes, err := userService.Login(getContext(c), req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	if req.RoleID == "34e9ec0f-aa20-42ab-b70d-4ba456c3bc48" {
		c.JSON(http.StatusOK, gin.H{
			"token": loginRes.Token,
			"role":  "customer",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"token": loginRes.Token,
			"role":  "admin",
		})
	}
}

func GetAccountByUsingAccessToken(c *context.Context) {
	c.JSON(http.StatusOK, c.User)
}

func GetAccessFromRefreshToken(c *gin.Context) {
	at, err := user.NewUsers().GetAccessFromRefreshToken(getContext(c), c.Query("refresh-token"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": at,
	})
}