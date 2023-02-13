package request

import (
	"getjv/backend/app/http/exception"
	"getjv/backend/app/http/models"
	"getjv/backend/kernel"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type AuthRequest struct{
		Username string `validate:"required,email"`
		Name string `validate:"required"`
		Password string `validate:"required"`
	}


func SignupPostRequest(c *fiber.Ctx) (models.User, exception.ErrorMessage) {
	
	validate := validator.New()
	authRequest := new(AuthRequest)
	
	if errParser := c.BodyParser(authRequest); errParser != nil {
		return models.User{},exception.ParserMessage(errParser)
	}
		
	errValidator := validate.Struct(authRequest)
	if errValidator != nil {
		return models.User{},exception.ValidationMessage(errValidator)
	}

	hashPass,errPassGen := kernel.GeneratePassword(authRequest.Password)
	if errPassGen != nil {
		return models.User{},exception.ServerMessage(errPassGen)
	}

	user := models.User{
		Username: authRequest.Username,
		Name: authRequest.Name,
		Password: string(hashPass),
	}

	return user,exception.ErrorMessage{}

}

