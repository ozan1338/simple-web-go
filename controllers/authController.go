package controllers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/ozan1338/simple-web/database"
	"github.com/ozan1338/simple-web/models"
	"github.com/ozan1338/simple-web/util"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	if data["password"] != data["password_confirm"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"status": "Error",
			"message": "Password Not Match",
		})
	}

	user := models.User{
		FirstName: data["first_name"],
		LastName: data["last_name"],
		Email: data["email"],
		RoleId: 3,
	}

	user.SetPassword(data["password"])

	go database.DB.Create(&user)

	return c.Status(201).JSON(fiber.Map{
		"status":"Success",
		"data": user,
	})
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "Error",
			"message": "Something Went Wrong",
		})
	}

	var user models.User

	// database.DB.Where("email = ?", data["email"]).First(&user)
	sqlQuery := "select * from users where email = ?"
	database.DB.Raw(sqlQuery, data["email"]).Scan(&user)

	//fmt.Println(user)

	if user.Id == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":"error",
			"message":"Not Found",
		})
	}

	if err := user.ComparePassword(data["password"]); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":"error",
			"message":"Incorect Password",
		})
	}

	token, err := util.GenerateJWT(strconv.Itoa(int(user.Id)))

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	cookie := fiber.Cookie{
		Name: "jwt",
		Value: token,
		Expires: time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.Status(200).JSON(fiber.Map{
		"status":"success",
		"data": user,
		"token":token,
	})
}

type Claims struct{
	jwt.StandardClaims
}

func User(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	id, _ := util.ParseJWT(cookie)

	var user models.User

	database.DB.Where("id = ?", id).First(&user)

	return c.Status(200).JSON(fiber.Map{
		"status":"success",
		"data":user,
	})

}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name: "jwt",
		Value: "",
		Expires: time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.Status(204).JSON(fiber.Map{
		"status":"success",
	})
}

func UpdateInfo(c *fiber.Ctx) error {
	var data map[string]string

	if err:= c.BodyParser(&data); err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	fmt.Println(data)
	cookie := c.Cookies("jwt")

	id,_ := util.ParseJWT(cookie)

	userId,_ := strconv.Atoi(id)

	user := models.User{
		Id: uint(userId),
		FirstName: data["first_name"],
		LastName: data["last_name"],
		Email: data["email"],
	}

	//fmt.Println(user)

	database.DB.Model(&user).Where("id: ?", id).Updates(user)

	return c.Status(200).JSON(fiber.Map{
		"status":"success",
		"data":user,
	})
}

func UpdatePassword(c *fiber.Ctx) error {
	var data map[string]string

	if err:= c.BodyParser(&data); err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	//fmt.Println(data)

	if data["password"] != data["password_confirm"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"status": "Error",
			"message": "Password Not Match",
		})
	}

	cookie := c.Cookies("jwt")

	id,_ := util.ParseJWT(cookie)

	user := models.User{}

	user.SetPassword(data["password"])

	database.DB.Model(&user).Where("id: ?", id).Updates(user)

	return c.Status(200).JSON(fiber.Map{
		"status":"success",
		"data":user,
	})
}