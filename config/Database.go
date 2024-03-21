package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

type DbConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	DbName   string
}

func BuildDbConfig() *DbConfig {
	dbConfig := DbConfig{
		Host:     "localhost",
		Port:     3306,
		Username: "root",
		Password: "",
		DbName:   "book_restapi",
	}
	return &dbConfig
}

func DbURL(dbConfig *DbConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DbName,
	)
}
