package domain

import (
	"database/sql"
	"errors"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql" //driver
	"github.com/jeffleon/banking-hexarch/errs"
)

type CustomerRepositoryDb struct {
	client *sql.DB
}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
	rows, err := QueryBuilderAllcustomers(status, d)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		}
		log.Println("Error while querying customer table" + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}
	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.ID, &c.Name, &c.City, &c.DateofBirth, &c.Zipcode, &c.Status)
		if err != nil {
			log.Println("Error while Scanning Customers" + err.Error())
			return nil, errs.NewUnexpectedError("error while scanning customers")
		}
		customers = append(customers, c)
	}
	return customers, nil
}

func (d CustomerRepositoryDb) ByID(id string) (*Customer, *errs.AppError) {
	customerSQL := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"
	row := d.client.QueryRow(customerSQL, id)
	var c Customer
	err := row.Scan(&c.ID, &c.Name, &c.City, &c.DateofBirth, &c.Zipcode, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		}
		log.Println("Error while querying customer table" + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}
	return &c, nil
}

func QueryBuilderAllcustomers(status string, d CustomerRepositoryDb) (*sql.Rows, error) {
	var findAllSQL string
	var rows *sql.Rows
	var err error
	if status == "" {
		findAllSQL = "select customer_id, name, city, zipcode, date_of_birth, status from customers"
		rows, err = d.client.Query(findAllSQL)
	} else {
		findAllSQL = "select customer_id, name, city, zipcode, date_of_birth, status from customers where status=?"
		if status == "active" {
			status = "1"
		} else if status == "inactive" {
			status = "0"
		} else {
			return nil, errors.New("unexpected status")
		}
		rows, err = d.client.Query(findAllSQL, status)
	}
	return rows, err
}

func NewCustomerRepositoryDB() CustomerRepositoryDb {
	client, err := sql.Open("mysql", "root:secret@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return CustomerRepositoryDb{client}
}
