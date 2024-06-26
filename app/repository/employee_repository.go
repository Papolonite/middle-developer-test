package repository

import (
	"middle-developer-test/app/model"

	"github.com/jmoiron/sqlx"
)

type EmployeeRepository struct {
	*sqlx.DB
}

func (repo *EmployeeRepository) CreateEmployee(employeeItem *model.Employee) (int, error) {
	query := `INSERT INTO employee (first_name, last_name, email, hire_date)
		VALUES	 ( $1, $2, $3, $4 ) RETURNING id;
	`

	var id int
	err := repo.QueryRow(query, employeeItem.FirstName, employeeItem.LastName, employeeItem.Email, employeeItem.HireDate).Scan(&id)
	if err != nil {
		return -1, err
	}

	return id, nil
}

func (repo *EmployeeRepository) GetAllEmployee() ([]model.Employee, error) {
	employeeList := []model.Employee{}

	query := "SELECT * FROM employee;"

	err := repo.Select(&employeeList, query)

	if err != nil {
		return nil, err
	}

	return employeeList, nil
}

func (repo *EmployeeRepository) GetEmployeeById(id int) (model.Employee, error) {
	employeeItem := model.Employee{}

	query := `SELECT * FROM employee WHERE id = $1;`

	err := repo.Get(&employeeItem, query, id)

	if err != nil {
		return employeeItem, err
	}

	return employeeItem, nil
}

func (repo *EmployeeRepository) UpdateEmployeeById(id int, employeeItem *model.Employee) error {
	query := `UPDATE employee set
		first_name = $2,
		last_name = $3,
		email = $4,
		hire_date = $5
		WHERE id = $1;
	`

	_, err := repo.Exec(query, id, employeeItem.FirstName, employeeItem.LastName, employeeItem.Email, employeeItem.HireDate)

	if err != nil {
		return err
	}

	return nil
}

func (repo *EmployeeRepository) DeleteEmployeeById(id int) error {
	query := `DELETE FROM employee WHERE id = $1`

	_, err := repo.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}
