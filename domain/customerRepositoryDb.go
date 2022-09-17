package domain

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" //driver
	"github.com/jeffleon/banking-lib/errs"
	"github.com/jeffleon/banking-lib/logger"
	"github.com/jmoiron/sqlx"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
	customers, err := QueryBuilderAllcustomers(status, d)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		}
		logger.Error("Error while querying customer table" + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}
	return customers, nil
}

func (d CustomerRepositoryDb) ByID(id string) (*Customer, *errs.AppError) {
	customerSQL := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"
	var c Customer
	err := d.client.Get(&c, customerSQL, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		}
		logger.Error("Error while querying customer table" + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}
	return &c, nil
}

func QueryBuilderAllcustomers(status string, d CustomerRepositoryDb) ([]Customer, error) {
	var findAllSQL string
	customers := make([]Customer, 0)
	var err error
	if status == "" {
		findAllSQL = "select customer_id, name, city, zipcode, date_of_birth, status from customers"
		err = d.client.Select(&customers, findAllSQL)
	} else {
		findAllSQL = "select customer_id, name, city, zipcode, date_of_birth, status from customers where status=?"
		err = d.client.Select(&customers, findAllSQL, status)
	}
	return customers, err
}

func NewCustomerRepositoryDB(client *sqlx.DB) CustomerRepositoryDb {
	return CustomerRepositoryDb{client}
}
