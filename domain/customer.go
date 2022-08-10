package domain

import (
	"github.com/jeffleon/banking-hexarch/dto"
	"github.com/jeffleon/banking-hexarch/errs"
)

type Customer struct {
	ID          string `db:"customer_id"`
	Name        string
	City        string
	Zipcode     string
	DateofBirth string `db:"date_of_birth"`
	Status      string
}

func (customer Customer) StatusAsText() string {
	StatusAsText := "active"
	if customer.Status == "0" {
		StatusAsText = "inactive"
	}
	return StatusAsText
}

func (customer Customer) ToDto() dto.CustomerResponse {
	response := dto.CustomerResponse{
		ID:          customer.ID,
		City:        customer.City,
		Name:        customer.Name,
		Zipcode:     customer.Zipcode,
		DateofBirth: customer.DateofBirth,
		Status:      customer.StatusAsText(),
	}
	return response
}

type CustomerRepository interface {
	FindAll(string) ([]Customer, *errs.AppError)
	ByID(string) (*Customer, *errs.AppError)
}
