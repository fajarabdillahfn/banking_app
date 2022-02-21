package domain

import "github.com/fajarabdillahfn/banking_app/errs"

type Customer struct {
	ID          string `db:"customer_id"`
	Name        string `db:"name"`
	City        string `db:"city"`
	Zipcode     string `db:"zipcode"`
	DateOfBirth string `db:"date_of_birth"`
	Status      string `db:"status"`
}

type CustomerRepository interface {
	FindAll(status string) ([]Customer, *errs.AppError)
	ByID(string) (*Customer, *errs.AppError)
}
