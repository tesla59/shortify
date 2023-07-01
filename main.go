package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
	"github.com/tesla59/shortify/database"
	"github.com/tesla59/shortify/handlers"
)

func main() {
	app := fiber.New(fiber.Config{})
	database.InitDB()
	handlers.InitRouter(app)
	app.Listen(os.Getenv("PORT"))
}
