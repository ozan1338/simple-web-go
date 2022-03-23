package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ozan1338/simple-web/controllers"
	"github.com/ozan1338/simple-web/middleware"
)

func Setup(app *fiber.App) {
    //AUTH
    app.Post("/api/register", controllers.Register)
    app.Post("/api/login", controllers.Login)

    app.Use(middleware.IsAuthenticated)

    app.Post("/api/logout", controllers.Logout)
    app.Get("/api/getUser", controllers.User)

    //CRUD USER
    app.Get("/api/users", controllers.AllUsers)
    app.Post("/api/users", controllers.CreateUser)
    app.Get("/api/users/:userId", controllers.GetUser)
    app.Patch("/api/users/:userId", controllers.UpdateUser)
    app.Delete("/api/users/:userId", controllers.DeleteUser)
}