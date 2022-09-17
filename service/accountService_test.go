package service

import (
	"testing"

	"github.com/golang/mock/gomock"
	realdomain "github.com/jeffleon/banking-hexarch/domain"
	"github.com/jeffleon/banking-hexarch/dto"
	"github.com/jeffleon/banking-hexarch/errs"
	"github.com/jeffleon/banking-hexarch/mocks/domain"
)

func Test_should_return_a_validation_error_response_when_the_request_is_not_validate(t *testing.T) {
	// Arrange
	req := dto.NewAccountRequest{
		CustomerID:  "100",
		AccountType: "saving",
		Amount:      0,
	}
	service := NewAccountService(nil)

	// Act
	_, appError := service.NewAccount(req)
	//
	if appError == nil {
		t.Error("failed while testing the new account validation")
	}

}

func Test_should_return_an_error_from_the_server_side_if_the_account_cannot_be_created(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := domain.NewMockAccountRepository(ctrl)
	service := NewAccountService(mockRepo)
	req := dto.NewAccountRequest{
		CustomerID:  "100",
		AccountType: "saving",
		Amount:      6000,
	}
	account := realdomain.Account{
		CustomerID:  req.CustomerID,
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      "1",
		OpeningDate: dbTSLayout,
	}
	mockRepo.EXPECT().Save(account).Return(nil, errs.NewUnexpectedError("unexpected"))
	// Act
	_, appError := service.NewAccount(req)
	// Assertion
	if appError == nil {
		t.Error("Test failed while validating error for new account")
	}
}
