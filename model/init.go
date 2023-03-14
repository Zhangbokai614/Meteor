package model

import (
	"fmt"

	"github.com/Zhangbokai614/go-template/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var conn *gorm.DB

const (
	userName = "root"
	password = "123456"
	host     = "10.0.1.182"
	port     = "3306"
	Database = "testing"
	charset  = "utf8mb4"

	adminUserName = "admin"
	AdminRoleName = "admin"
	adminRouter   = "admin"
)

var (
	dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=%s&parseTime=True&loc=Local",
		userName, password, host, port, charset)
)

func GetDBConnection() *gorm.DB {
	mysqlDB := mysql.Open(dsn)

	conn, err := gorm.Open(mysqlDB, &gorm.Config{})
	if err != nil {
		panic(err)
	}
	conn.Exec("USE " + Database)

	return conn
}

func init() {
	var (
		err error
	)

	mysqlDB := mysql.Open(dsn)
	conn, err = gorm.Open(mysqlDB, &gorm.Config{})
	if err != nil {
		panic(err)
	}

	conn.Exec("CREATE DATABASE IF NOT EXISTS " + Database + " DEFAULT CHARACTER SET " + charset)
	conn.Exec("USE " + Database)

	conn.AutoMigrate(&User{})
	conn.AutoMigrate(&Role{})
	conn.AutoMigrate(&RolePermissions{})
	conn.AutoMigrate(&Permissions{})

	userPermissions := []*Permissions{
		{RouterPermissions: adminRouter},
		{RouterPermissions: "/api/v1/user/create"},
		{RouterPermissions: "/api/v1/permissions/create/role"},
		{RouterPermissions: "/api/v1/permissions/query/role"},
		{RouterPermissions: "/api/v1/permissions/modify/role/permissions"},
	}
	conn.Clauses(clause.Insert{Modifier: "IGNORE"}).Create(userPermissions)

	role := &Role{
		Name: AdminRoleName,
	}
	conn.Clauses(clause.Insert{Modifier: "IGNORE"}).Create(role)

	rp := &RolePermissions{
		PID: userPermissions[0].ID,
		RID: role.ID,
	}
	conn.Clauses(clause.Insert{Modifier: "IGNORE"}).Create(&rp)

	conn.Clauses(clause.Insert{Modifier: "IGNORE"}).Create(&User{Name: adminUserName, RID: role.ID, Password: utils.Md5Encode("123456")})
}
