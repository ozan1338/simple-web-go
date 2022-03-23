package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ozan1338/simple-web/database"
	"github.com/ozan1338/simple-web/models"
)

func AllUsers(c *fiber.Ctx) error {
	var users []models.User

	database.DB.Find(&users)

	return c.JSON(fiber.Map{
		"status":"success",
		"result":len(users),
		"data":users,
	})
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	user.SetPassword("1234")

	database.DB.Create(&user)

	return c.Status(201).JSON(fiber.Map{
		"status":"success",
		"data":user,
	})
}