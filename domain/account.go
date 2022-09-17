package domain

import (
	"github.com/jeffleon/banking-hexarch/dto"
	"github.com/jeffleon/banking-lib/errs"
)

type Account struct {
	AccountID   string
	CustomerID  string
	OpeningDate string
	AccountType string
	Amount      float64
	Status      string
}

func (a Account) ToNewAccountResponseDto() dto.NewAccountResponse {
	return dto.NewAccountResponse{AccountID: a.AccountID}
}

//go:generate mockgen -destination=../mocks/domain/mockAccountRepository.go -package=domain github.com/jeffleon/banking-hexarch/domain AccountRepository
type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
}
