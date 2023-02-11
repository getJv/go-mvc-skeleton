package resource

import (
	"getjv/backend/app/http/models"

	"github.com/gofiber/fiber/v2"
)

func UserCollection(c *fiber.Ctx, users []models.User) error {
    
	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"users": users,
	})
}