package controller

import (
	"getjv/backend/app/http/exception"
	"getjv/backend/app/http/request"
	"getjv/backend/app/http/resource/users"
	"getjv/backend/app/http/service"

	"github.com/gofiber/fiber/v2"
)

func UserIndex(c *fiber.Ctx) error {
	
	// Get all users.
	users, serviceErr := service.UserFindAll(c)
	if serviceErr.Error{
		return exception.ErrorResponse(c,serviceErr)
	}
	
	return resource.UserCollection(c,users)
}

func UserStore(c *fiber.Ctx) error  {
	
	userRequest,validationErr := request.UserPostRequest(c)
	if validationErr.Error {
	   return exception.ErrorResponse(c,validationErr)
	}

	user := userRequest.ParseModel()

	_, serviceErr := service.UserAdd(c,&user)
	if serviceErr.Error {
		return exception.ErrorResponse(c,serviceErr)
	}
	
	return resource.UserResource(c,user);
}

func UserShow(c *fiber.Ctx) error  {
	id := c.Params("id")
	
	user, serviceErr := service.UserFind(c,id)
	if serviceErr.Error{
		return exception.ErrorResponse(c,serviceErr)
	}

	return resource.UserResource(c,user);
}

func UserDelete(c *fiber.Ctx) error  {
	id := c.Params("id")
			
	user, serviceFindErr := service.UserFind(c,id)
	if serviceFindErr.Error{
		return exception.ErrorResponse(c,serviceFindErr)
	}
	
	_, serviceDelErr := service.UserDelete(c,&user)
	if serviceDelErr.Error{
		return exception.ErrorResponse(c,serviceDelErr)
	}

	return resource.UserResource(c,user);
}

func UserUpdate(c *fiber.Ctx) error  {

	userRequest,validationErr := request.UserPostRequest(c)
	if validationErr.Error {
	   return exception.ErrorResponse(c,validationErr)
	}
	
	user, serviceFindErr := service.UserFind(c,userRequest.ID)
	if serviceFindErr.Error{
		return exception.ErrorResponse(c,serviceFindErr)
	}

	userRequest.Bind(&user)

	_, serviceUptErr := service.UserUpdate(c,&user)
	if serviceUptErr.Error{
		return exception.ErrorResponse(c,serviceUptErr)
	}

	return resource.UserResource(c,user);

}