package routes

import (
	"middle-developer-test/app/controller"

	"github.com/gofiber/fiber/v2"
)

func EmployeeRoutes(a *fiber.App) {
	route := a.Group("/api/employee")

	route.Get("/", controller.GetAllEmployee)
	route.Get("/:employeeId", controller.GetEmployeeById)
	route.Post("/", controller.CreateEmployee)
	route.Put("/:employeeId", controller.UpdateEmployeeById)
	route.Delete("/:employeeId", controller.DeleteEmployee)
}
