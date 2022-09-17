package service

import (
	"github.com/jeffleon/banking-hexarch/domain"
	"github.com/jeffleon/banking-hexarch/dto"
	"github.com/jeffleon/banking-lib/errs"
)

type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
}

type DefaultAccountService struct {
	repo domain.AccountRepository
}

const dbTSLayout = "2006-01-02 15:04:05"

func (s DefaultAccountService) NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	a := domain.Account{
		AccountID:   "",
		CustomerID:  req.CustomerID,
		OpeningDate: dbTSLayout,
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      "1",
	}
	newAccount, err := s.repo.Save(a)
	if err != nil {
		return nil, err
	}
	responseDto := newAccount.ToNewAccountResponseDto()
	return &responseDto, nil
}

func NewAccountService(repo domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo: repo}
}
