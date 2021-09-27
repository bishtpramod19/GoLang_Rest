package dto

import "Banking/errs"

const WITHDRAWAL = "withdrawl"
const DEPOSIT = "deposit"

type TransactionRequest struct {
	TransactionType string  `json:"transaction_type"`
	Amount          float64 `json:"amount"`
	AccountId       string  `json:"account_id"`
	TransactionDate string  `json:"transaction_date"`
	CustomerId      string  `json:"-"`
}

func (r TransactionRequest) IsTransactionTypeWithdrawl() bool {
	return r.TransactionType == WITHDRAWAL
}

func (r TransactionRequest) IsTransactionTypeDeposit() bool {
	return r.TransactionType == DEPOSIT
}

func (r TransactionRequest) Validate() *errs.AppError {
	if !r.IsTransactionTypeWithdrawl() && !r.IsTransactionTypeDeposit() {
		return errs.NewValidationError("Transaction type can only be either Withdrawl or Deposit !")
	}
	if r.Amount < 0 {
		return errs.NewValidationError("Amount cannot be less than zero")
	}

	return nil

}

type TransactionResponse struct {
	TransactionId   string  `jsom:"transaction_id"`
	AccountId       string  `json:"account_id"`
	Amount          float64 `json:"new_balance"`
	TransactionType string  `json:"transaction_type"`
	TransactionDate string  `json:"transaction_date"`
}
