package middlewares

import (
	"errors"
	"net/http"

	"github.com/Zhangbokai614/go-template/model"
	"github.com/gin-gonic/gin"
)

var (
	errorPermissionsVerifyFailed = errors.New("User permissions verify failed")
	errorPermissionsInsufficient = errors.New("User insufficient permissions")
)

func RbacPermissionsVerify() gin.HandlerFunc {
	dbConn := model.GetDBConnection()

	return func(c *gin.Context) {
		rid := c.GetUint(ContextKeyUserRID)

		role := &model.Role{
			ID: rid,
		}
		if err := dbConn.Model(role).Select("name").Where(role).First(&role).Error; err != nil {
			c.JSON(http.StatusInternalServerError, errorPermissionsVerifyFailed.Error())
			c.Abort()

			return
		}
		if role.Name == model.AdminRoleName {
			c.Next()

			return
		}

		rp := &model.RolePermissions{
			RID: rid,
		}

		pids := make([]uint, 0)
		if err := dbConn.Model(rp).Select("p_id").Where(rp).Scan(&pids).Error; err != nil || len(pids) == 0 {
			c.JSON(http.StatusInternalServerError, errorPermissionsVerifyFailed.Error())
			c.Abort()

			return
		}

		p := &model.Permissions{
			RouterPermissions: c.Request.RequestURI,
		}

		if r := dbConn.Model(p).Select("id").Where("id in ?", pids).Where(p).RowsAffected; r == 0 {
			c.JSON(http.StatusInternalServerError, errorPermissionsInsufficient.Error())
			c.Abort()

			return
		}

		c.Next()
	}
}
