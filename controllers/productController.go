package controllers

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/ozan1338/simple-web/database"
	"github.com/ozan1338/simple-web/models"
)

func AllProducts(c *fiber.Ctx) error {
	limit := 5
	page,_ := strconv.Atoi( c.Query("page","1"))
	offset := (page-1) * limit
	var products []models.Product

	database.DB.Limit(limit).Offset(offset).Find(&products)

	return c.JSON(fiber.Map{
		"status":"success",
		"result":len(products),
		"data":products,
	})
}

func CreateProduct(c *fiber.Ctx) error {
	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		fmt.Println(err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	database.DB.Create(&product)

	return c.Status(201).JSON(fiber.Map{
		"status":"success",
		"data":product,
	})
}

func GetProduct(c *fiber.Ctx) error {
	productId,_ := strconv.Atoi(c.Params("productId"))

	Product := models.Product{
		Id: uint(productId),
	}

	database.DB.Find(&Product)

	if Product.Id == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":"error",
			"message":"Not Found",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":"success",
		"data":Product,
	})
}

func UpdateProduct(c *fiber.Ctx) error {
	productId,_ := strconv.Atoi(c.Params("productId"))

	product := models.Product{
		Id: uint(productId),
	}

	if err := c.BodyParser(&product); err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	//database.DB.Model(&user).Updates(user)
	sqlQuery := models.QueryUpdateProduct(product)

	database.DB.Raw(sqlQuery, product.Id).Scan(&product)

	return c.Status(200).JSON(fiber.Map{
		"status":"success",
		"data": product,
	})
}

func DeleteProduct(c *fiber.Ctx) error {
	productId,_ := strconv.Atoi(c.Params("productId"))

	product := models.Product{
		Id: uint(productId),
	}

	database.DB.Delete(&product)

	return c.SendStatus(fiber.StatusNoContent)
}