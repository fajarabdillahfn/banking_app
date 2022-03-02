package domain

import (
	"log"
	"strconv"

	"github.com/fajarabdillahfn/banking-lib/errs"
	"github.com/fajarabdillahfn/banking-lib/logger"
	"github.com/jmoiron/sqlx"
)

type AccountRepositoryDb struct {
	client *sqlx.DB
}

func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{client: dbClient}
}

func (d AccountRepositoryDb) Save(a Account) (*Account, *errs.AppError) {
	var id int64

	sqlInsert := `INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) 
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

func (d AccountRepositoryDb) FindAccount(id string) (*Account, *errs.AppError) {
	var a Account

	customerSql := `SELECT * 
					FROM accounts 
					WHERE account_id = $1`

	err := d.client.Get(&a, customerSql, id)

	if err != nil {
		logger.Error("Error while scanning customer: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	return &a, nil
}

func (d AccountRepositoryDb) SaveTransaction(t Transaction) (*Transaction, *errs.AppError) {
	var updateSql string
	var id int64

	tx, err := d.client.Begin()
	if err != nil {
		logger.Error("Error while starting a new transaction for bank account transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	transactionSql := `INSERT INTO transactions (account_id, amount, transaction_type, transaction_date) 
					   VALUES ($1, $2, $3, $4)
					   RETURNING transaction_id`

	stmt, err := d.client.Prepare(transactionSql)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	err = stmt.QueryRow(t.AccountID, t.Amount, t.TransactionType, t.TransactionDate).Scan(&id)
	if err != nil {
		logger.Error("Error while getting last transaction account id: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	if t.IsWithdrawal() {
		updateSql = `UPDATE accounts SET amount = amount - $1 WHERE account_id=$2`
	} else {
		updateSql = `UPDATE accounts SET amount = amount + $1 WHERE account_id=$2`
	}

	_, err = tx.Exec(updateSql, t.Amount, t.AccountID)
	if err != nil {
		tx.Rollback()
		logger.Error("Error while saving transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		logger.Error("Error while commiting transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	account, appErr := d.FindAccount(t.AccountID)
	if appErr != nil {
		return nil, appErr
	}

	t.TransactionID = strconv.FormatInt(id, 10)
	t.Amount = account.Amount

	return &t, nil
}
