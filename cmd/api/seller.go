package main

import (
	"context"
	"time"
	"github.com/rejath-chandran/go-ecom/internal/data"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	
)

func (app *App) Login(c *fiber.Ctx) error {

	claims := jwt.MapClaims{
		"name":  "John Doe",
		"admin": true,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte("rejathchandran"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})
}

func (app *App) Details(c *fiber.Ctx) error {

	

	data:=data.User{
		Name: "rejath",
		Email: "rejath13@gmail.com",
	}

	coll := app.Con.Collection("Users")
	_, err := coll.InsertOne(context.TODO(), data)
	if err != nil {
		panic(err)
	}
	name, ok := c.Locals("username").(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid username in context"})
	}

	return c.SendString("Welcome " + name)
}

func GetToken(c *fiber.Ctx) error {

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	c.Locals("username", name)

	return c.Next()
}
