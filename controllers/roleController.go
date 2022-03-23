package controllers

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/ozan1338/simple-web/database"
	"github.com/ozan1338/simple-web/models"
)

func AllRoles(c *fiber.Ctx) error {
	var roles []models.Role

	database.DB.Find(&roles)

	return c.JSON(fiber.Map{
		"status":"success",
		"result":len(roles),
		"data":roles,
	})
}

func CreateRole(c *fiber.Ctx) error {
	var role models.Role

	if err := c.BodyParser(&role); err != nil {
		fmt.Println(err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	database.DB.Create(&role)

	return c.Status(201).JSON(fiber.Map{
		"status":"success",
		"data":role,
	})
}

func GetRole(c *fiber.Ctx) error {
	roleId,_ := strconv.Atoi(c.Params("roleId"))

	role := models.Role{
		Id: uint(roleId),
	}

	database.DB.Find(&role)

	if role.Name == ""  {
		return c.Status(404).JSON(fiber.Map{
			"status":"error",
			"message":"Not Found",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":"success",
		"data":role,
	})
}

func UpdateRole(c *fiber.Ctx) error {
	roleId,_ := strconv.Atoi(c.Params("roleId"))

	role := models.Role{
		Id: uint(roleId),
	}

	if err := c.BodyParser(&role); err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	sqlQuery := models.QueryUpdateRole(role)

	database.DB.Raw(sqlQuery, role.Id).Scan(&role)
	//database.DB.Model(&role).Updates(role)

	return c.Status(200).JSON(fiber.Map{
		"status":"success",
		"data": role,
	})
}

func DeleteRole(c *fiber.Ctx) error {
	roleId,_ := strconv.Atoi(c.Params("roleId"))

	role := models.Role{
		Id: uint(roleId),
	}

	fmt.Println("Nihao")

	//database.DB.Delete(&role)
	sqlQuery := models.QueryUpdateRoleIsActive(false)
	database.DB.Raw(sqlQuery, role.Id).Scan(&role)

	return c.SendStatus(fiber.StatusNoContent)
}