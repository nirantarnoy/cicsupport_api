package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	dbUserMysql = "cicuser"
	dbPassMysql = "Ax12345678"
	// dbUserMysql = "root"
	// dbPassMysql = ""

	db = [2]string{"192.168.60.189", "cicsupport"}

	dsn1 = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUserMysql, dbPassMysql,
		db[0], db[1])
)

func SetupDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn1), &gorm.Config{})
	if err != nil {
		panic("Failed to Connect Database")
	}

	return db
}

func CloseDBConn(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close Database.")
	}
	dbSQL.Close()
}
