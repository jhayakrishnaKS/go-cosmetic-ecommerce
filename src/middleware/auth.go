package middlewares

import (
	"ecommerce/src/constants"
	"ecommerce/src/services/user"
	"ecommerce/src/utils/context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func WithAuth(next func(*context.Context)) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := GetBearerToken(c)
		ctx := getContext(c)

		

		user, err := user.NewUsers().GetAccountWithAccessToken(ctx, token)
		if err == constants.ErrAccessTokenExpired {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.User = userFromRequest(user)
		next(ctx)
	}
}
