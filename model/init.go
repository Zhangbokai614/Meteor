package model

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var conn *gorm.DB

const (
	userName = "root"
	password = "123456"
	host     = "10.0.1.182"
	port     = "3306"
	Database = "user"
	charset  = "utf8mb4"
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

	db, err := conn.DB()
	if err != nil {
		panic(err)
	}

	db.SetConnMaxIdleTime(0)
	db.SetConnMaxLifetime(0)

	conn.Exec("CREATE DATABASE IF NOT EXISTS " + Database + " DEFAULT CHARACTER SET " + charset)

	conn.AutoMigrate(&User{})
}
