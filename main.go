// 🚀 Fiber is an Express inspired web framework written in Go with 💖
// 📌 API Documentation: https://docs.gofiber.io
// 📝 Github Repository: https://github.com/gofiber/fiber
package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/n4sunday/go-fiber-mongo/database"
	"github.com/n4sunday/go-fiber-mongo/routes"
)

func main() {
	// Connect to the database
	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()
	routes.SetupRoute(app)
	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON("Hello Go Fiber MongoDB 🚀")
	})
	log.Fatal(app.Listen(":3000"))
}
