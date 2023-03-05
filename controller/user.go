package controller

import (
	"net/http"

	"github.com/Zhangbokai614/go-template/middlewares"
	"github.com/Zhangbokai614/go-template/model"
	"github.com/Zhangbokai614/go-template/service"
	"github.com/Zhangbokai614/go-template/utils"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	data := &model.APIUser{}

	if err := c.BindJSON(&data); err != nil {
		middlewares.LogError(c, err)
		c.JSON(http.StatusBadRequest, &gin.H{
			"message": "Invalid Parameter",
		})

		return
	}

	if service.UserNameIsExists(data.Name) {
		c.JSON(StatusUserError, &gin.H{
			"message": ErrorUserAlreadyExists.Error(),
		})

		return
	}

	if err := service.CreateUser(data); err != nil {
		middlewares.LogError(c, err)
		c.JSON(http.StatusInternalServerError, &gin.H{
			"message": err,
		})

		return
	}

	c.JSON(http.StatusOK, &gin.H{
		"message": "User create success",
	})
	return
}

func UserLogin(c *gin.Context) {
	type user struct {
		Name     string `json:"name" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	data := &user{}
	if err := c.BindJSON(&data); err != nil {
		middlewares.LogError(c, err)
		c.JSON(http.StatusBadRequest, &gin.H{
			"message": "Invalid Parameter",
		})

		return
	}

	if !service.UserNameIsExists(data.Name) {
		c.JSON(StatusUserError, &gin.H{
			"message": ErrorUserDoesNotExists.Error(),
		})

		return
	}

	result, err := service.QueryUserByName(data.Name)
	if err != nil {
		middlewares.LogError(c, err)
		c.JSON(StatusUserError, &gin.H{
			"message": "Username or Password incorrect",
		})

		return
	}

	if !utils.Md5Check(data.Password, result.Password) {
		c.JSON(StatusUserError, &gin.H{
			"message": "Username or Password incorrect",
		})

		return
	}

	token, err := utils.GenerateToken(result.ID)
	if err != nil {
		middlewares.LogError(c, err)
		c.JSON(http.StatusForbidden, &gin.H{"message": "Failed to apply for token"})

		return
	}

	c.JSON(http.StatusOK, &gin.H{"token": token})
}
