package middlewares

import "github.com/gin-gonic/gin"

const (
	ContextKeyTraceID = "Custom-Trace-ID"
)

var (
	traceID int64
)

func init() {
	traceID = 1000000
}

func AttachTraceID() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceID += 1
		c.Set(ContextKeyTraceID, traceID)

		c.Next()
	}
}
