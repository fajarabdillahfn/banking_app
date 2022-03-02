package domain

import (
	"github.com/fajarabdillahfn/banking-lib/errs"
	"github.com/fajarabdillahfn/banking_app/dto"
)

type Customer struct {
	ID          string `db:"customer_id"`
	Name        string `db:"name"`
	City        string `db:"city"`
	Zipcode     string `db:"zipcode"`
	DateOfBirth string `db:"date_of_birth"`
	Status      string `db:"status"`
}

func (c Customer) statusAsText() string {
	if c.Status == "0" {
		return "inactive"
	} else if c.Status == "1" {
		return "active"
	}
	return ""
}

func (c Customer) ToDto() dto.CustomerResponse {
	return dto.CustomerResponse{
		ID:          c.ID,
		Name:        c.Name,
		City:        c.City,
		Zipcode:     c.Zipcode,
		DateOfBirth: c.DateOfBirth,
		Status:      c.statusAsText(),
	}
}

type CustomerRepository interface {
	FindAll(status string) ([]Customer, *errs.AppError)
	ByID(string) (*Customer, *errs.AppError)
}
