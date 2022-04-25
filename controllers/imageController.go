package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func Upload(c *fiber.Ctx) error {
	form, err := c.MultipartForm()

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":"Error",
			"message":err,
		})
	}

	files := form.File["image"]
	filename := ""

	for _, item := range files {
		filename = item.Filename

		if err := c.SaveFile(item, "./upload/" + filename); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"status":"Error",
				"message":err,
			})
		}
	}

	return c.Status(201).JSON(fiber.Map{
		"status":"success",
		"url":"http://localhost:8000/api/upload/" + filename,
	})
}