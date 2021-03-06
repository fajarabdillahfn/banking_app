package domain

import (
	"github.com/fajarabdillahfn/banking-lib/errs"
	"github.com/fajarabdillahfn/banking_app/dto"
)

type Account struct {
	AccountID   string  `db:"account_id"`
	CustomerID  string  `db:"customer_id"`
	OpeningDate string  `db:"opening_date"`
	AccountType string  `db:"account_type"`
	Amount      float64 `db:"amount"`
	Status      string  `db:"status"`
}

type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
	FindAccount(string) (*Account, *errs.AppError)
	SaveTransaction(Transaction) (*Transaction, *errs.AppError)
}

func (a Account) ToNewAccountResponseDto() dto.NewAccountResponse {
	return dto.NewAccountResponse{AccountID: a.AccountID}
}

func (a Account) CanWithdraw(amount float64) bool {
	return a.Amount >= amount
}
