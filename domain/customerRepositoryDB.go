package domain

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type CustomerRepositoryDb struct {
	client *sql.DB
}

func (d CustomerRepositoryDb) FindAll() ([]Customer, error) {
	findAllSql := "SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers"

	rows, err := d.client.Query(findAllSql)
	if err != nil {
		log.Println("Error while querying table:", err)
		return nil, err
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
			return nil, err
		}
		customers = append(customers, c)
	}

	return customers, nil
}

func (d CustomerRepositoryDb) ByID(id string) (*Customer, error) {
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
		log.Println("Error while scanning customer " + err.Error())
		return nil, err
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
