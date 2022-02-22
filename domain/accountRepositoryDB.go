package domain

import (
	"log"
	"strconv"

	"github.com/fajarabdillahfn/banking_app/errs"
	"github.com/fajarabdillahfn/banking_app/logger"
	"github.com/jmoiron/sqlx"
)

type AccountRepositoryDb struct {
	client *sqlx.DB
}

func (d AccountRepositoryDb) Save(a Account) (*Account, *errs.AppError) {
	var id int64

	sqlInsert := `INSERT INTO accounts(customer_id, opening_date, account_type, amount, status) 
				  VALUES ($1, $2, $3, $4, $5)
				  RETURNING account_id`

	stmt, err := d.client.Prepare(sqlInsert)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	err = stmt.QueryRow(a.CustomerID, a.OpeningDate, a.AccountType, a.Amount, a.Status).Scan(&id)
	if err != nil {
		logger.Error("Error while creating new account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	a.AccountID = strconv.FormatInt(id, 10)
	return &a, nil
}

func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{client: dbClient}
}
