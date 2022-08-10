package domain

import "github.com/jeffleon/banking-hexarch/errs"

type Customer struct {
	ID          string `db:"customer_id"`
	Name        string
	City        string
	Zipcode     string
	DateofBirth string `db:"date_of_birth"`
	Status      string
}

type CustomerRepository interface {
	FindAll(string) ([]Customer, *errs.AppError)
	ByID(string) (*Customer, *errs.AppError)
}
