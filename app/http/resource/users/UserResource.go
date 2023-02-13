package resource

import (
	"getjv/backend/app/http/models"

	"github.com/gofiber/fiber/v2"
)

func UserResource(c *fiber.Ctx, user models.User ) error {
	user.Password = ""
	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"user": user,
	})
}