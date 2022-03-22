package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ozan1338/simple-web/controllers"
)

func Setup(app *fiber.App) {
    app.Post("/api/register", controllers.Register)
    app.Post("/api/login", controllers.Login)
    app.Get("/api/getUser", controllers.User)
    app.Post("/api/logout", controllers.Logout)
}