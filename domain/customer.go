package domain

import "github.com/fajarabdillahfn/banking_app/app/errs"

type Customer struct {
	ID          string
	Name        string
	City        string
	Zipcode     string
	DateOfBirth string
	Status      string
}

type CustomerRepository interface {
	FindAll() ([]Customer, *errs.AppError)
	ByID(string) (*Customer, *errs.AppError)
}
