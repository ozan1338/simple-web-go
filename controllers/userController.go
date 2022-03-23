package controllers

import (
	"strconv"

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

func GetUser(c *fiber.Ctx) error {
	userId,_ := strconv.Atoi(c.Params("userId"))

	user := models.User{
		Id: uint(userId),
	}

	database.DB.Find(&user)

	if user.FirstName == "" || user.LastName == "" || user.Email == "" {
		return c.Status(404).JSON(fiber.Map{
			"status":"error",
			"message":"Not Found",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":"success",
		"data":user,
	})
}

func UpdateUser(c *fiber.Ctx) error {
	userId,_ := strconv.Atoi(c.Params("userId"))

	user := models.User{
		Id: uint(userId),
	}

	if err := c.BodyParser(&user); err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	//database.DB.Model(&user).Updates(user)
	sqlQuery := models.QueryUpdateUser(user)

	database.DB.Raw(sqlQuery, user.Id).Scan(&user)

	return c.Status(200).JSON(fiber.Map{
		"status":"success",
		"data": user,
	})
}

func DeleteUser(c *fiber.Ctx) error {
	userId,_ := strconv.Atoi(c.Params("userId"))

	user := models.User{
		Id: uint(userId),
	}

	database.DB.Delete(&user)

	return c.SendStatus(fiber.StatusNoContent)
}