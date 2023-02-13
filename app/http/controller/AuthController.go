package controller

import (
	"getjv/backend/app/http/exception"
	"getjv/backend/app/http/request"
	resource "getjv/backend/app/http/resource/users"
	"getjv/backend/app/http/service"
	"getjv/backend/kernel"

	"github.com/gofiber/fiber/v2"
)

func AuthSignUp(c *fiber.Ctx) error  {
	
	user,validationErr := request.SignupPostRequest(c)
	if validationErr.Error {
	   return exception.ErrorResponse(c,validationErr)
	}
 	
	_, serviceErr := service.UserAdd(c,&user)
	if serviceErr.Error {
		return exception.ErrorResponse(c,serviceErr)
	}
	
	return resource.UserResource(c,user);
}

// AuthLogin Do Login.
// @Description Do Login.
// @Summary Do Login.
// @Tags Auth
// @Param payload body string true "Json message" SchemaExample({\r\n    "username":"jhonatan@test.com",\r\n    "password":"123456"\r\n})
// @Param        id   path      int  true  "Account ID"
// @Accept json
// @Produce json
// @Success 200 {object}	models.User
// @Failure 401 {object}	exception.ErrorMessage
// @Failure 503 {object}	exception.ErrorMessage
// @Router /api/login [post]
func AuthLogin(c *fiber.Ctx) error  {
	
	userRequest,validationErr := request.AuthPostLoginRequest(c)
	if validationErr.Error {
	   return exception.ErrorResponse(c,validationErr)
	}
 	
	localUser, serviceErr := service.UserFindByEmail(c,userRequest.Username)
	if serviceErr.Error {
		return exception.ErrorResponse(c,serviceErr)
	}

	authErr := kernel.ComparePassword(localUser.Password,userRequest.Password)
	if authErr != nil{
		return exception.ErrorResponse(c,exception.UnAuthMessage())
	}

	token,errToken := kernel.GenerateToken(localUser)
	if errToken != nil{
		return exception.ErrorResponse(c,exception.ServerMessage(errToken))
	}

	errCookie := kernel.SaveAuthCookie(c,token)
	if errCookie != nil{
		return exception.ErrorResponse(c,exception.ServerMessage(errCookie))
	}
	
	return resource.AuthResource(c,localUser,token);
}


func AuthValidate(c *fiber.Ctx) error  {
	
	authUser,errAuthUser := kernel.GetAuthUser(c)
	if errAuthUser != nil{
		return exception.ErrorResponse(c,exception.UnAuthMessage())
	}
	
	return resource.UserResource(c, authUser)
}




