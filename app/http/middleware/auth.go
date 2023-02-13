package middleware

import (
	"getjv/backend/app/http/exception"
	"getjv/backend/app/http/service"
	"getjv/backend/kernel"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func RequireAuth(c *fiber.Ctx) error{
	
	userId,errSession := kernel.RetrieveUserFromCookie(c)
	if errSession != nil || userId == 0 {
		return exception.ErrorResponse(c,exception.UnAuthMessage())
	}

	user, serviceErr := service.UserFind(c, strconv.Itoa(userId))
	if serviceErr.Error {
		return exception.ErrorResponse(c,serviceErr)
	}
	
	c.Locals("user", user)
	
	return c.Next()
}