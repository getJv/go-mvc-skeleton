package resource

import (
	"getjv/backend/app/http/models"

	"github.com/gofiber/fiber/v2"
)

func AuthResource(c *fiber.Ctx,user models.User,token string ) error {
    
	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"user": user,
		"token": token,
	})
}