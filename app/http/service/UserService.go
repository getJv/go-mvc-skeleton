package service

import (
	"getjv/backend/app/http/exception"
	"getjv/backend/app/http/models"
	"getjv/backend/kernel"

	"github.com/gofiber/fiber/v2"
)

func UserFindAll(c *fiber.Ctx ) ( []models.User,exception.ErrorMessage)  {
	var users []models.User
	result := kernel.DB.Find(&users)
	if result.Error != nil{
		return []models.User{},exception.DatabaseMessage(result.Error)
	}
	return users,exception.ErrorMessage{}
}

func UserFind(c *fiber.Ctx,id string ) ( models.User,exception.ErrorMessage)  {
	var user models.User		
	result := kernel.DB.Find(&user,id)
	if result.Error != nil{
		return models.User{},exception.DatabaseMessage(result.Error)
	}
	if result.RowsAffected == 0{
		return models.User{},exception.NotFoundMessage()
	}
	return user,exception.ErrorMessage{}
}

func UserAdd(c *fiber.Ctx, user *models.User) ( *models.User ,exception.ErrorMessage)  {
	result := kernel.DB.Create(&user)
	if result.Error != nil{
		c.Status(500)
		return &models.User{},exception.DatabaseMessage(result.Error)
	}
	return user,exception.ErrorMessage{}
}

func UserDelete(c *fiber.Ctx, user *models.User ) ( models.User,exception.ErrorMessage)  {
			
	result := kernel.DB.Delete(user)
	if result.Error != nil{
		return models.User{},exception.DatabaseMessage(result.Error)
	}
	return models.User{},exception.ErrorMessage{}
}

func UserUpdate(c *fiber.Ctx, user *models.User ) ( models.User,exception.ErrorMessage)  {
			
	result := kernel.DB.Model(&user).Updates(user)
	if result.Error != nil{
		return models.User{},exception.DatabaseMessage(result.Error)
	}
	return *user,exception.ErrorMessage{}
}

