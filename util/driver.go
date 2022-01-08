package util

import (
	"fmt"
	"layerarch/config"
	"layerarch/entities"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(config *config.AppConfig) *gorm.DB {

	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", 
		config.Database.Username,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.Name,
	)

	db, err := gorm.Open(mysql.Open(conn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db
}

func InitialMigrate(db *gorm.DB)  {
	db.AutoMigrate(&entities.User{})
	db.AutoMigrate(&entities.Book{})
}