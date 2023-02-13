package models

import (
	"gorm.io/gorm"
)

type User struct {
  gorm.Model 
  Username  string `json:"username" example:"username@test.com"`
  Name string `json:"name" example:"Ada"`
  Password string `json:"password" swaggerignore:"true"`
}



