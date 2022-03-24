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
	var roleDetail fiber.Map

	if err := c.BodyParser(&roleDetail); err != nil {
		fmt.Println(err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	list := roleDetail["permission"].([]interface{})

	permission := make([]models.Permission, len(list))

	for i, permissionId := range list {
		id , _ := strconv.Atoi(permissionId.(string))

		permission[i] = models.Permission{
			Id: uint(id),
		}
	}

	role := models.Role{
		Name: roleDetail["name"].(string),
		Permission: permission,
	}

	role.IsActive = true

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

	database.DB.Preload("Permission").Find(&role)

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

	var roleDetail fiber.Map

	if err := c.BodyParser(&roleDetail); err != nil {
		fmt.Println(err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	list := roleDetail["permission"].([]interface{})

	permission := make([]models.Permission, len(list))

	for i, permissionId := range list {
		id , _ := strconv.Atoi(permissionId.(string))

		permission[i] = models.Permission{
			Id: uint(id),
		}
	}

	role := models.Role{
		Id: uint(roleId),
		Name: roleDetail["name"].(string),
		Permission: permission,
	}
	
	if len(role.Permission) > 0 {
		var result interface{}
		sqlQueryDeletePermission := models.QueryDeleteDeletePermission()
		database.DB.Raw(sqlQueryDeletePermission, role.Id).Scan(result)
		for i:=0 ; i < len(role.Permission); i++ {
			sqlQueryForUpdatePermission := models.QueryUpdateRolePermission(role.Permission[i], role.Id)
			database.DB.Raw(sqlQueryForUpdatePermission).Scan(result)
			fmt.Println(sqlQueryForUpdatePermission)
		}
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