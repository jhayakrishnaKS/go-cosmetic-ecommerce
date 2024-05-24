package middlewares

import (
	"ecommerce/src/database/models"
	"ecommerce/src/dtos"
	"ecommerce/src/utils/context"
	"ecommerce/src/utils/db"

	"strings"

	"github.com/gin-gonic/gin"
)

func getContext(c *gin.Context) *context.Context {
	return &context.Context{
		Context: c,
		DB:      db.New(),
	}
}

func GetBearerToken(c *gin.Context) string {
	bearerToken := c.Request.Header.Get("Authorization")
	if bearerToken != "" {
		strs := strings.Split(bearerToken, " ")
		if len(strs) > 1 && strs[0] == "Bearer" {
			return strs[1]
		}
	}
	return ""
}

func userFromRequest(req *models.Users) *dtos.User {
	return &dtos.User{
		ID:       req.ID,
		Email:    req.Email,
		Username: req.Username,
		RoleID:   req.RoleID,
	}
}
