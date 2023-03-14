package routers

import (
	"github.com/Zhangbokai614/go-template/controller"
	"github.com/Zhangbokai614/go-template/middlewares"
	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {
	router.POST("/ping", controller.Ping)

	v1 := router.Group("/api/v1")
	v1.POST("/user/login", controller.UserLogin)

	user := v1.Group("/user", middlewares.RbacPermissionsVerify())
	user.Use(middlewares.JwtAuthMiddleware(), middlewares.RbacPermissionsVerify())
	user.POST("/create", controller.CreateUser)

	permissions := v1.Group("/permissions")
	permissions.Use(middlewares.JwtAuthMiddleware(), middlewares.RbacPermissionsVerify())
	permissions.POST("/create/role", controller.CreateRole)
	permissions.POST("/query/role", controller.QueryRole)
	permissions.POST("/query/permissions", controller.QueryPermissions)
	permissions.POST("/delete/role", controller.DeleteRole)
	permissions.POST("/modify/role/permissions", controller.ModifyRolePermissions)
	permissions.POST("/modify/user/role", controller.ModifyUserRole)
}
