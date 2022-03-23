package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ozan1338/simple-web/controllers"
	"github.com/ozan1338/simple-web/middleware"
)

func Setup(app *fiber.App) {
    app.Post("/api/register", controllers.Register)
    app.Post("/api/login", controllers.Login)

    app.Use(middleware.IsAuthenticated)

    app.Post("/api/logout", controllers.Logout)
    app.Get("/api/getUser", controllers.User)
    app.Get("/api/users", controllers.AllUsers)
    app.Post("/api/users", controllers.CreateUser)
}