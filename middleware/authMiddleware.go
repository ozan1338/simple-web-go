package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ozan1338/simple-web/util"
)

func IsAuthenticated(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	if _, err := util.ParseJWT(cookie); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":"Error",
			"message":"Unauthorized",
		})

	}

	return c.Next()
}