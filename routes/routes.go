package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/n4sunday/go-fiber-mongo/modules/employee"
	"github.com/n4sunday/go-fiber-mongo/modules/position"
	"github.com/n4sunday/go-fiber-mongo/modules/switchs"
)

func SetupRoute(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/employee", employee.GetAllEmployee)
	api.Get("/employee/:id", employee.GetEmployee)
	api.Post("/employee", employee.CreateEmployee)
	api.Put("/employee/:id", employee.UpdateEmployee)
	api.Delete("/employee/:id", employee.DeleteEmployee)

	api.Get("/position", position.GetAllPosition)
	api.Get("/position/:id", position.GetPosition)
	api.Post("/position", position.CreatePosition)
	api.Put("/position/:id", position.UpdatePosition)
	api.Delete("/position/:id", position.DeletePosition)

	api.Get("/switch", switchs.GetAllSwitch)
}
