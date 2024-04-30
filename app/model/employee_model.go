package model

import (
	"time"
)

type Employee struct {
	Id        int       `db:"id" json:"id"`
	FirstName string    `db:"first_name" json:"firstName"`
	LastName  string    `db:"last_name" json:"lastName"`
	Email     string    `db:"email" json:"email"`
	HireDate  time.Time `db:"hire_date" json:"hireDate"`
}
