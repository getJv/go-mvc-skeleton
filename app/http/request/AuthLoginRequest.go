package request

import (
	"getjv/backend/app/http/exception"
	"getjv/backend/app/http/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type AuthLoginRequest struct{
		Username string `json:"username" validate:"required,email" example:"jhonatan@test.com"`
		Password string `json:"password" validate:"required" example:"123456"`
	}


func AuthPostLoginRequest(c *fiber.Ctx) (models.User, exception.ErrorMessage) {
	
	validate := validator.New()
	authRequest := new(AuthLoginRequest)
	
	if errParser := c.BodyParser(authRequest); errParser != nil {
		return models.User{},exception.ParserMessage(errParser)
	}
		
	errValidator := validate.Struct(authRequest)
	if errValidator != nil {
		return models.User{},exception.ValidationMessage(errValidator)
	}

	user := models.User{
		Username: authRequest.Username,
		Password: authRequest.Password,
	}

	return user,exception.ErrorMessage{}

}

