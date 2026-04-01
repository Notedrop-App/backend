package main

import (
	"os"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v3"
	"github.com/notedrop-app/backend/utils"
)

func init() {
	utils.LoadENV("./.env")
}

func main() {

	port := os.Getenv("PORT")

	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello World!")
	})

	app.Listen(":" + port)
}
