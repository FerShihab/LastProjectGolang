package configs

import (
	"lastProject/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	dsn := "root:Clayjr54!@tcp(127.0.0.1:3306)/prakerja8?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	initMigration()
}

func initMigration() {

	DB.AutoMigrate(&models.Car{})
}
