package kernel

import (
	"getjv/backend/app/http/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func DatabaseConnect() *gorm.DB{
  	
	db, err := gorm.Open(sqlite.Open("database/database.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.Logger = logger.Default.LogMode(logger.Info)
	return db
	  
}

func LoadMigrations(){
	DB.AutoMigrate(&models.User{})
}