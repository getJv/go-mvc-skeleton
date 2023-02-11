package request

import (
	"getjv/backend/app/http/exception"
	"getjv/backend/app/http/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type UserRequest struct{
		ID string
		Username string `validate:"required"`
		Name string `validate:"required"`
	}

// Useful onCreate operation
func (userRequest *UserRequest) ParseModel() models.User{
	return models.User{
		Username: userRequest.Username,
		Name: userRequest.Name,
	}
}

// Useful onUpdate operation
func (userRequest *UserRequest) Bind(user *models.User){
	if(user.Username != userRequest.Username){
    	user.Username = userRequest.Username
  	}

  	if(user.Name != userRequest.Name){
    	user.Name = userRequest.Name
  	}
} 



func UserPostRequest(c *fiber.Ctx) (UserRequest, exception.ErrorMessage) {
	
	validate := validator.New()
	userRequest := new(UserRequest)
	
	if errParser := c.BodyParser(userRequest); errParser != nil {
		return UserRequest{},exception.ParserMessage(errParser)
	}
		
	errValidator := validate.Struct(userRequest)
	if errValidator != nil {
		return UserRequest{},exception.ValidationMessage(errValidator)
	}

	user := UserRequest{
		ID: c.Params("id"),
		Username: userRequest.Username,
		Name: userRequest.Name,
	}

	return user,exception.ErrorMessage{}

}