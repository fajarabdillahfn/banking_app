package dto

import (
	"strings"

	"github.com/fajarabdillahfn/banking_app/errs"
)

type TransactionRequest struct {
	CustomerID      string  `json:"customer_id"`
	AccountID       string  `json:"account_id"`
	TransactionType string  `json:"transaction_type"`
	Amount          float64 `json:"amount"`
}

type TransactionResponse struct {
	TransactionID   string  `json:"transaction_id"`
	AccountID       string  `json:"account_id"`
	NewBalance      float64 `json:"new_balance"`
	TransactionType string  `json:"transaction_type"`
	TransactionDate string  `json:"transaction_date"`
}

func (r TransactionRequest) Validate() *errs.AppError {
	if r.Amount < 0 {
		return errs.NewValidationError("Transaction cannot be negative amount")
	}

	if strings.ToLower(r.TransactionType) != "withdraw" && strings.ToLower(r.TransactionType) != "deposit" {
		return errs.NewValidationError("Transaction type should be 'withdraw' or 'deposit'")
	}
	return nil
}

func (r TransactionRequest) IsTransactionWithdrawal() bool {
	return strings.ToLower(r.TransactionType) == "withdraw"
}
