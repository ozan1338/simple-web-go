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

    app.Patch("/api/users/info", controllers.UpdateInfo)
    app.Patch("/api/users/password", controllers.UpdatePassword)
    app.Post("/api/logout", controllers.Logout)
    app.Get("/api/getUser", controllers.User)

    //CRUD USER
    app.Get("/api/users", controllers.AllUsers)
    app.Post("/api/users", controllers.CreateUser)
    app.Get("/api/users/:userId", controllers.GetUser)
    app.Patch("/api/users/:userId", controllers.UpdateUser)
    app.Delete("/api/users/:userId", controllers.DeleteUser)

    //CRUD ROLE
    app.Get("/api/role", controllers.AllRoles)
    app.Post("/api/role", controllers.CreateRole)
    app.Get("/api/role/:roleId", controllers.GetRole)
    app.Patch("/api/role/:roleId", controllers.UpdateRole)
    app.Delete("/api/role/:roleId", controllers.DeleteRole)

    //CRUD PERMISSION
    app.Get("/api/permission", controllers.AllPermission)

    //CRUD PRODUCT
    app.Get("/api/product", controllers.AllProducts)
    app.Post("/api/product", controllers.CreateProduct)
    app.Get("/api/product/:productId", controllers.GetProduct)
    app.Patch("/api/product/:productId", controllers.UpdateProduct)
    app.Delete("/api/product/:productId", controllers.DeleteProduct)

    app.Post("/api/upload", controllers.Upload)
    app.Static("/api/upload", "./upload")
}