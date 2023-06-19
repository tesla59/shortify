package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tesla59/shortify/database"
	"github.com/tesla59/shortify/handlers"
)

func main() {
	app := fiber.New(fiber.Config{})
	database.InitDB()
	handlers.InitRouter(app)
	app.Listen(":5566")
}
