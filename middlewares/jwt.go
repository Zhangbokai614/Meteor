package middlewares

import (
	"net/http"

	"github.com/Zhangbokai614/go-template/utils"
	"github.com/gin-gonic/gin"
)

const (
	ContextKeyUserID   = "custom-user-id"
	ContextKeyUserUnit = "custom-user-unit"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := utils.TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "message: Token invalid")
			c.Abort()
			return
		}

		id, err := utils.ExtractTokenID(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "message: Extract token name fail")
			c.Abort()
			return
		}

		c.Set(ContextKeyUserID, id)

		c.Next()
	}
}
