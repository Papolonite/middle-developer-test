package database

import (
	"fmt"
	"middle-developer-test/app/repository"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

type Repositories struct {
	*repository.EmployeeRepository
}

func PostgreSQLConection() (*Repositories, error) {
	db, err := sqlx.Connect("pgx", "postgresql://dev:dev@localhost:5102/postgres?sslmode=disable")
	if err != nil {
		return nil, fmt.Errorf("cannot connect to database, %w", err)
	}

	return &Repositories{
		EmployeeRepository: &repository.EmployeeRepository{DB: db},
	}, nil
}
