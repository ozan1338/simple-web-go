package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ozan1338/simple-web/database"
	"github.com/ozan1338/simple-web/models"
)

func AllPermission(c *fiber.Ctx) error {
	var permission []models.Permission

	database.DB.Find(&permission)

	return c.JSON(fiber.Map{
		"status": "success",
		"result": len(permission),
		"data":   permission,
	})
}