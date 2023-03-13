package middlewares

import (
	"fmt"
	"net/http"

	"github.com/Zhangbokai614/go-template/utils"
	"github.com/gin-gonic/gin"
)

const (
	ContextKeyUserID  = "custom-user-id"
	ContextKeyUserRID = "custom-user-rid"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := utils.TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "message: Token invalid")
			c.Abort()
			return
		}

		id, rid, err := utils.ExtractTokenID(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "message: Extract token name fail")
			c.Abort()
			return
		}

		fmt.Println("jwt------", id, rid)
		c.Set(ContextKeyUserID, id)
		c.Set(ContextKeyUserRID, rid)

		c.Next()
	}
}
