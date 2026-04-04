package main

import (
	"fmt"
	"os"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v3"
	"github.com/notedrop-app/backend/db"
	"github.com/notedrop-app/backend/utils"
)

func init() {
	utils.LoadENV("./.env")
	err := db.ConnectDB(os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Println("Error connecting to database", err)
	}
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
