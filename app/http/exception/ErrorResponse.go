package exception

import (
	"github.com/gofiber/fiber/v2"
)

func ErrorResponse(c *fiber.Ctx, errMsg ErrorMessage ) error {
	
	return c.Status(errMsg.Code).JSON(errMsg)
}