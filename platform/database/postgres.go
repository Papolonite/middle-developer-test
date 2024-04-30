package database

import (
	"fmt"
	"middle-developer-test/app/repository"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

type Repositories struct {
	*repository.EmployeeRepository
}

func PostgreSQLConection() (*Repositories, error) {
	db, err := sqlx.Connect("pgx", os.Getenv("DB_URL"))
	if err != nil {
		return nil, fmt.Errorf("cannot connect to database, %w", err)
	}

	return &Repositories{
		EmployeeRepository: &repository.EmployeeRepository{DB: db},
	}, nil
}
