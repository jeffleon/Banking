package domain

import "github.com/jeffleon/banking-hexarch/errs"

type Customer struct {
	ID          string
	Name        string
	City        string
	Zipcode     string
	DateofBirth string
	Status      string
}

type CustomerRepository interface {
	FindAll() ([]Customer, error)
	ByID(string) (*Customer, *errs.AppError)
}
