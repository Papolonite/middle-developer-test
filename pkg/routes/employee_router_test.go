package routes

import (
	"encoding/json"
	"io"
	"middle-developer-test/app/model"
	"middle-developer-test/pkg/config"
	"middle-developer-test/platform/database"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

type GetEmployeeListResponse struct {
	Message string           `json:"message"`
	Data    []model.Employee `json:"data"`
}

func TestEmployeeRoutes(t *testing.T) {
	setUpTest()
	defer TearDownTest()

	app := fiber.New()
	EmployeeRoutes(app)

	// Create Employee Test
	postEmployeeTest := []struct {
		description        string
		route              string
		requestBody        io.Reader
		expectedStatusCode int
		expectedError      bool
	}{
		{
			description: "create new employee",
			route:       "/api/employee",
			requestBody: strings.NewReader(`{
				"first_name" : "Test",
				"last_name" : "Employee",
				"email" : "testEmployee@gmail.com",
				"hire_date" : "2024-04-23T00:00:00.000Z"
		}`),
			expectedStatusCode: 200,
			expectedError:      false,
		},
		{
			description: "create new employee with uncomplete body",
			route:       "/api/employee",
			requestBody: strings.NewReader(`{
				"first_name" : "Test2",
				"last_name" : "Employee",
		}`),
			expectedStatusCode: 400,
			expectedError:      false,
		},
		{
			description: "create new employee with wrong email format",
			route:       "/api/employee",
			requestBody: strings.NewReader(`{
				"first_name" : "Test3",
				"last_name" : "Employee",
				"email" : "testEmployee",
				"hire_date" : "2024-04-23T00:00:00.000Z"
		}`),
			expectedStatusCode: 400,
			expectedError:      false,
		},
	}
	for _, test := range postEmployeeTest {
		req := httptest.NewRequest("POST", test.route, test.requestBody)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req, -1)

		assert.Equalf(t, test.expectedError, err != nil, test.description)
		if test.expectedError {
			continue
		}

		assert.Equalf(t, test.expectedStatusCode, resp.StatusCode, test.description)
	}

	// Get All Employee List
	getAllEmployeeTest := []struct {
		description        string
		route              string
		expectedStatusCode int
		expectedError      bool
		expectedLength     int
	}{
		{
			description:        "get all employee",
			route:              "/api/employee",
			expectedStatusCode: 200,
			expectedError:      false,
			expectedLength:     1,
		},
	}

	var getAllEmployeeResponse GetEmployeeListResponse
	var firstItemId int

	for _, test := range getAllEmployeeTest {
		req := httptest.NewRequest("GET", test.route, nil)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req, -1)

		assert.Equalf(t, test.expectedError, err != nil, test.description)
		if test.expectedError {
			continue
		}

		assert.Equalf(t, test.expectedStatusCode, resp.StatusCode, test.description)

		readResponse, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		err = json.Unmarshal(readResponse, &getAllEmployeeResponse)
		if err != nil {
			panic(err)
		}

		assert.Equalf(t, test.expectedLength, len(getAllEmployeeResponse.Data), test.description)
		firstItemId = getAllEmployeeResponse.Data[0].Id
	}

	// Get Employee List
	getEmployeeTest := []struct {
		description        string
		route              string
		expectedStatusCode int
		expectedError      bool
	}{
		{
			description:        "get employee",
			route:              "/api/employee/" + strconv.Itoa(firstItemId),
			expectedStatusCode: 200,
			expectedError:      false,
		},
		{
			description:        "get nonexistent employee",
			route:              "/api/employee/" + strconv.Itoa(firstItemId+len(getAllEmployeeResponse.Data)),
			expectedStatusCode: 404,
			expectedError:      false,
		},
	}

	for _, test := range getEmployeeTest {
		req := httptest.NewRequest("GET", test.route, nil)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req, -1)

		assert.Equalf(t, test.expectedError, err != nil, test.description)
		if test.expectedError {
			continue
		}

		assert.Equalf(t, test.expectedStatusCode, resp.StatusCode, test.description)
	}

	// Update Employee Test
	updateEmployeeTest := []struct {
		description        string
		route              string
		requestBody        io.Reader
		expectedStatusCode int
		expectedError      bool
	}{
		{
			description: "update employee",
			route:       "/api/employee/" + strconv.Itoa(firstItemId),
			requestBody: strings.NewReader(`{
				"first_name" : "TestUpdated",
				"last_name" : "EmployeeUpdated"
		}`),
			expectedStatusCode: 200,
			expectedError:      false,
		},
		{
			description: "update non existent employee",
			route:       "/api/employee/" + strconv.Itoa(firstItemId+len(getAllEmployeeResponse.Data)),
			requestBody: strings.NewReader(`{
				"first_name" : "TestUpdated",
				"last_name" : "EmployeeUpdated"
		}`),
			expectedStatusCode: 404,
			expectedError:      false,
		},
		{
			description: "update employee",
			route:       "/api/employee/" + strconv.Itoa(firstItemId),
			requestBody: strings.NewReader(`{
				"email" : "TestUpdated",
		}`),
			expectedStatusCode: 400,
			expectedError:      false,
		},
	}

	for _, test := range updateEmployeeTest {
		req := httptest.NewRequest("PUT", test.route, test.requestBody)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req, -1)

		assert.Equalf(t, test.expectedError, err != nil, test.description)
		if test.expectedError {
			continue
		}

		assert.Equalf(t, test.expectedStatusCode, resp.StatusCode, test.description)
	}

	// Delete Employee List
	deleteEmployeeTest := []struct {
		description        string
		route              string
		expectedStatusCode int
		expectedError      bool
	}{
		{
			description:        "delete employee",
			route:              "/api/employee/" + strconv.Itoa(firstItemId),
			expectedStatusCode: 200,
			expectedError:      false,
		},
		{
			description:        "delete nonexistent employee",
			route:              "/api/employee/" + strconv.Itoa(firstItemId+len(getAllEmployeeResponse.Data)),
			expectedStatusCode: 404,
			expectedError:      false,
		},
	}

	for _, test := range deleteEmployeeTest {
		req := httptest.NewRequest("DELETE", test.route, nil)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req, -1)

		assert.Equalf(t, test.expectedError, err != nil, test.description)
		if test.expectedError {
			continue
		}

		assert.Equalf(t, test.expectedStatusCode, resp.StatusCode, test.description)
	}
}

func setUpTest() {
	config.LoadConfig("../../.env.local.test")
}

func TearDownTest() {
	db, err := database.PostgreSQLConection()
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("TRUNCATE TABLE employee CASCADE;")
	if err != nil {
		panic(err)
	}
}
