package dto

import (
	"strings"

	"github.com/jeffleon/banking-hexarch/errs"
)

type NewAccountRequest struct {
	CustomerID  string  `json:"customer_id"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}

func (a NewAccountRequest) Validate() *errs.AppError {
	if a.Amount < 5000 {
		return errs.NewValidationError("To open a new account you need to deposite at least 5000.00")
	}
	if strings.ToLower(a.AccountType) != "saving" && strings.ToLower(a.AccountType) != "checking" {
		return errs.NewValidationError("Account type should be cheking or saving")
	}
	return nil
}
