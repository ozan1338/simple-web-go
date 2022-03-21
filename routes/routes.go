package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ozan1338/simple-web/controllers"
)

func Setup(app *fiber.App) {
    app.Get("/", controllers.Hello)
}