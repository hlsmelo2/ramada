package db

import (
	"fmt"
	"ramada/api/src/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func getDB(host string, port string, dbname string, user string, password string) (*gorm.DB, error) {
	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user,
		password,
		host,
		port,
		dbname,
	)

	return gorm.Open(mysql.Open(connectionString), &gorm.Config{})
}

func GetDB() *gorm.DB {
	db, _error := getDB(
		config.DB_HOST,
		config.DB_PORT,
		config.DB_NAME,
		config.DB_USER,
		config.DB_PASSWORD,
	)

	if _error != nil {
		panic("failed to connect database")
	}

	return db
}
