package controller

import (
	"errors"
	"net/http"

	"github.com/Zhangbokai614/go-template/middlewares"
	"github.com/Zhangbokai614/go-template/model"
	"github.com/Zhangbokai614/go-template/service"
	"github.com/gin-gonic/gin"
)

var (
	errorRoleAlreadyExists = errors.New("Role already exists")
	errorRoleNotExists     = errors.New("Role not exists")
)

func CreateRole(c *gin.Context) {
	data := model.APIRole{}

	if err := c.BindJSON(&data); err != nil {
		middlewares.LogError(c, err)
		c.JSON(http.StatusBadRequest, &gin.H{
			"message": "Invalid Parameter",
		})

		return
	}

	if service.RoleNameIsExists(data.Name) {
		c.JSON(http.StatusBadRequest, &gin.H{
			"message": errorRoleAlreadyExists.Error(),
		})

		return
	}

	if err := service.CreateRole(&data); err != nil {
		c.JSON(http.StatusBadRequest, &gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, &gin.H{
		"message": "Create role success",
	})

	return
}

func QueryRole(c *gin.Context) {
	roles, err := service.QueryRole()
	if err != nil {
		c.JSON(http.StatusBadRequest, &gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, roles)

	return
}

func QueryPermissions(c *gin.Context) {
	roles, err := service.QueryPermissions()
	if err != nil {
		c.JSON(http.StatusBadRequest, &gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, roles)

	return
}

func DeleteRole(c *gin.Context) {
	var data struct {
		RID uint `json:"r_id"`
	}

	if err := c.BindJSON(&data); err != nil {
		middlewares.LogError(c, err)
		c.JSON(http.StatusBadRequest, &gin.H{
			"message": "Invalid Parameter",
		})

		return
	}

	err := service.DeleteRole(data.RID)
	if err != nil {
		c.JSON(http.StatusBadRequest, &gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, &gin.H{
		"message": "Delete success",
	})

	return
}

func ModifyRolePermissions(c *gin.Context) {
	data := model.APIRole{}

	if err := c.BindJSON(&data); err != nil {
		middlewares.LogError(c, err)
		c.JSON(http.StatusBadRequest, &gin.H{
			"message": "Invalid Parameter",
		})

		return
	}

	if !service.RoleNameIsExists(data.Name) {
		c.JSON(http.StatusBadRequest, &gin.H{
			"message": errorRoleNotExists.Error(),
		})

		return
	}

	if err := service.ModifyRolePermissions(&data); err != nil {
		c.JSON(http.StatusBadRequest, &gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, &gin.H{
		"message": "Create role success",
	})

	return
}
