package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	c := fiber.New()

	c.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Fatal(c.Listen(":" + port))
}
