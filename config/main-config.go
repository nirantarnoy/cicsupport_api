package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	dbUserMysql = "root"
	dbPassMysql = "Tamakogi2012"

	db = [2]string{"172.16.0.27", "cicsupport"}

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
