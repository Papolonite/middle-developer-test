package employee

import (
	"context"
	"time"

	"github.com/Papolonite/middle-developer-test/models"
	"github.com/jackc/pgx"
	"github.com/rs/zerolog/log"
)

func createEmployee(ctx context.Context, tx pgx.Tx, employeeItem models.Employee) error {
	query := `INSERT INTO employee (first_name, last_name, email, hire_date)
		VALUES	 ( $1, $2, $3, $4 );
	`

	_, err := tx.Exec(ctx, query, employeeItem.FirstName, employeeItem.LastName, employeeItem.Email, employeeItem.HireDate)

	if err != nil {
		return err
	}

	return nil
}

func getAllEmployee(ctx context.Context, tx pgx.Tx) ([]models.Employee, error) {
	var employeeCount int

	countRow := tx.QueryRow(ctx, "SELECT count(id) FROM employee;")
	err := countRow.Scan(&employeeCount)

	if err != nil {
		log.Warn().Err(err).Msg("cannot count employee table")
		return nil, err
	}

	if employeeCount == 0 {
		return nil, err
	}

	employeeItems := make([]models.Employee, employeeCount)

	rows, err := tx.Query(ctx, "ELECT * FROM employee;")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for i := range rows {
		var id int
		var firstName string
		var lastName string
		var email string
		var hireDate time.Time

		if !rows.Next() {
			break
		}

		if err := rows.Scan(&id, &firstName, &lastName, &email, &hireDate); err != nil {
			log.Warn().Err(err).Msg("cannot scan employee item")
			return nil, err
		}

		employeeItems[i] = models.Employee{
			Id:        id,
			FirstName: firstName,
			LastName:  lastName,
			Email:     email,
			HireDate:  hireDate,
		}
	}

	return employeeItems, nil
}

func getEmployeeById(ctx context.Context, tx pgx.Tx, id int) (models.Employee, error) {
	query := `SELECT * FROM employee WHERE id = $1;`
	rows := tx.QueryRow(ctx, query, id)

	var employeeItem models.Employee

	if err := rows.Scan(&employeeItem.Id, &employeeItem.FirstName, &employeeItem.LastName, &employeeItem.Email, &employeeItem.HireDate); err != nil {
		if err == pgx.ErrNoRows {
			log.Debug().Err(err).Msg("can't find employee based on id given")
		}
		return models.Employee{}, err
	}

	return employeeItem, nil
}

func updateEmployeeById(ctx context.Context, tx pgx.Tx, employeeItem models.Employee) error {
	query := `UPDATE employee set
		first_name = $1 ,
		last_name = $2 ,
		email = $3 ,
		hire_date = $4
		WHERE id = $5;
	`

	_, err := tx.Exec(ctx, query, employeeItem.FirstName, employeeItem.LastName, employeeItem.Email, employeeItem.HireDate, employeeItem.Id)

	if err != nil {
		return err
	}

	return nil
}

func deleteEmployeeById(ctx context.Context, tx pgx.Tx, id int) error {
	query := `DELETE employee WHERE id = $5;`

	_, err := tx.Exec(ctx, query, employeeItem.Id)

	if err != nil {
		return err
	}

	return nil
}
