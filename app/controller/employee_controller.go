package controller

import (
	"middle-developer-test/app/model"
	"middle-developer-test/pkg/lib"
	"middle-developer-test/platform/database"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func CreateEmployee(c *fiber.Ctx) error {
	employeeBody := &model.Employee{}

	if err := c.BodyParser(employeeBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"message": err.Error(),
			})
	}

	db, err := database.PostgreSQLConection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"message": err.Error(),
			})
	}

	if err := lib.ValidateEmail(employeeBody.Email); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"message": err.Error(),
			})
	}

	if err := db.CreateEmployee(employeeBody); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"message": err.Error(),
			})
	}

	return c.JSON(fiber.Map{
		"message": "Successfully created new employee",
	})
}

func GetAllEmployee(c *fiber.Ctx) error {
	db, err := database.PostgreSQLConection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"message": err.Error(),
			})
	}

	employeeList, err := db.GetAllEmployee()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"message": err.Error(),
			})
	}

	return c.JSON(fiber.Map{
		"message": "Successfully fetch all employee data",
		"data":    employeeList,
	})
}

func GetEmployeeById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("employeeId"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"message": err.Error(),
			})
	}

	db, err := database.PostgreSQLConection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"message": err.Error(),
			})
	}

	employee, err := db.GetEmployeeById(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(
			fiber.Map{
				"message": err.Error(),
			})
	}

	return c.JSON(fiber.Map{
		"message": "Successfully fetched employee data",
		"data":    employee,
	})
}

func UpdateEmployeeById(c *fiber.Ctx) error {
	// Check Request Body Valid
	employeeBody := &model.Employee{}

	if err := c.BodyParser(employeeBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"message": err.Error(),
			})
	}

	db, err := database.PostgreSQLConection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"message": err.Error(),
			})
	}

	if err := lib.ValidateEmail(employeeBody.Email); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"message": err.Error(),
			})
	}

	// Check If Employee Exist
	id, err := strconv.Atoi(c.Params("employeeId"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"message": err.Error(),
			})
	}

	db, err = database.PostgreSQLConection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"message": err.Error(),
			})
	}

	employee, err := db.GetEmployeeById(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(
			fiber.Map{
				"message": err.Error(),
			})
	}

	// Update Employee
	if employeeBody.FirstName == "" {
		employeeBody.FirstName = employee.FirstName
	}
	if employeeBody.LastName == "" {
		employeeBody.LastName = employee.LastName
	}
	if employeeBody.Email == "" {
		employeeBody.Email = employee.Email
	}
	if employeeBody.HireDate.IsZero() {
		employeeBody.HireDate = employee.HireDate
	}

	if err := db.UpdateEmployeeById(id, employeeBody); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"message": err.Error(),
			})
	}

	return c.JSON(fiber.Map{
		"message": "Successfully updated employee data",
	})
}

func DeleteEmployee(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("employeeId"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"message": err.Error(),
			})
	}

	// Check Employee Exist
	db, err := database.PostgreSQLConection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"message": err.Error(),
			})
	}

	employee, err := db.GetEmployeeById(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(
			fiber.Map{
				"message": err.Error(),
			})
	}

	// Delete Employee
	err = db.DeleteEmployeeById(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"message": err.Error(),
			})
	}

	return c.JSON(fiber.Map{
		"message": "Successfully deleted employee data"
	})
}
