package domain

import (
	"strings"

	"github.com/fajarabdillahfn/banking-lib/errs"
	"github.com/fajarabdillahfn/banking_app/dto"
)

type Transaction struct {
	TransactionID   string  `db:"transaction_id"`
	AccountID       string  `db:"account_id"`
	Amount          float64 `db:"amount"`
	TransactionType string  `db:"transaction_type"`
	TransactionDate string  `db:"transaction_date"`
}

type TransactionRepository interface {
	Withdraw(Transaction) (*Transaction, *errs.AppError)
}

func (t Transaction) ToNewTransactionResponseDto() dto.TransactionResponse {
	return dto.TransactionResponse{
		TransactionID:   t.TransactionID,
		AccountID:       t.AccountID,
		NewBalance:      t.Amount,
		TransactionType: t.TransactionType,
		TransactionDate: t.TransactionDate,
	}
}

func (t Transaction) IsWithdrawal() bool {
	return strings.ToLower(t.TransactionType) == "withdraw"
}
