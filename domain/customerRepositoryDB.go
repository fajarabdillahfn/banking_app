package domain

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/fajarabdillahfn/banking_app/errs"
	"github.com/fajarabdillahfn/banking_app/logger"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
	where := ""
	customers := make([]Customer, 0)

	if status != "" {
		where = fmt.Sprintf("WHERE status = %s", status)
	}

	findAllSql := fmt.Sprintf("SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers %s", where)

	err := d.client.Select(&customers, findAllSql)
	if err != nil {
		logger.Error("Error while querying table: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	return customers, nil
}

func (d CustomerRepositoryDb) ByID(id string) (*Customer, *errs.AppError) {
	var c Customer

	customerSql := `SELECT customer_id, name, city, zipcode, date_of_birth, status 
					FROM customers 
					WHERE customer_id = $1`

	err := d.client.Get(&c, customerSql, id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer Not Found")
		} else {
			logger.Error("Error while scanning customer " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}

	return &c, nil
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	dbDriver := os.Getenv("DB_DRIVER")
	dbUser := os.Getenv("DB_USER")
	dbPasswd := os.Getenv("DB_PASSWD")
	dbAddr := os.Getenv("DB_ADDR")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dataSource := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable", dbDriver, dbUser, dbPasswd, dbAddr, dbPort, dbName)
	client, err := sqlx.Open(dbDriver, dataSource)
	if err != nil {
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return CustomerRepositoryDb{client: client}
}
