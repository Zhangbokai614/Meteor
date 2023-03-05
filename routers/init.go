package routers

import (
	"github.com/Zhangbokai614/go-template/controller"
	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {
	router.POST("/ping", controller.Ping)

	v1 := router.Group("/api/v1")
	v1.POST("/user/create", controller.CreateUser)
	v1.POST("/user/login", controller.UserLogin)
}
