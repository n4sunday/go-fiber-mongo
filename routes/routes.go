package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/n4sunday/go-fiber-mongo/modules/employee"
)

func SetupRoute(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/employee", employee.GetAllEmployee)
	api.Get("/employee/:id", employee.GetEmployee)
	api.Post("/employee", employee.CreateEmployee)
	api.Put("/employee/:id", employee.UpdateEmployee)
	api.Delete("/employee/:id", employee.DeleteEmployee)
}
