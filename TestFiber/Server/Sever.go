package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/", func(req *fiber.Ctx) error {
		return req.SendString("Hello, World")
	})
	app.Listen(":3000")
}
