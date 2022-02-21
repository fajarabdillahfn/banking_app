package domain

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/fajarabdillahfn/banking_app/errs"
	_ "github.com/lib/pq"
)

type CustomerRepositoryDb struct {
	client *sql.DB
}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
	where := ""

	if status != "" {
		where = fmt.Sprintf("WHERE status = %s", status)
	}

	findAllSql := fmt.Sprintf("SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers %s", where)

	rows, err := d.client.Query(findAllSql)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		} else {
			log.Println("Error while querying table:", err)
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}

	customers := make([]Customer, 0)

	for rows.Next() {
		var c Customer
		err := rows.Scan(
			&c.ID,
			&c.Name,
			&c.City,
			&c.Zipcode,
			&c.DateOfBirth,
			&c.Status,
		)
		if err != nil {
			log.Println("Error while scanning:", err)
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
		customers = append(customers, c)
	}

	return customers, nil
}

func (d CustomerRepositoryDb) ByID(id string) (*Customer, *errs.AppError) {
	customerSql := `SELECT customer_id, name, city, zipcode, date_of_birth, status 
					FROM customers 
					WHERE customer_id = $1`

	row := d.client.QueryRow(customerSql, id)

	var c Customer

	err := row.Scan(
		&c.ID,
		&c.Name,
		&c.City,
		&c.Zipcode,
		&c.DateOfBirth,
		&c.Status,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer Not Found")
		} else {
			log.Println("Error while scanning customer " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}

	return &c, nil
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	client, err := sql.Open("postgres", "postgres://abdillah.fajar:masBed0311@localhost/banking_app?sslmode=disable")
	if err != nil {
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return CustomerRepositoryDb{client: client}
}
