package models

import (
	"time"
)

type Employee struct {
	Id        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	HireDate  time.Time `json:"hireDate"`
}
