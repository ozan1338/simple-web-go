package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ozan1338/simple-web/database"
	"github.com/ozan1338/simple-web/routes"
)

func main() {
	database.Connect()

	app := fiber.New()

	routes.Setup(app)

    app.Listen(":8000")
}