package app

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/jeffleon/banking-hexarch/domain"
	"github.com/jeffleon/banking-hexarch/mocks/service"
	"github.com/jeffleon/banking-lib/errs"
)

var router *mux.Router
var ch CustomerHandler

func setup(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	mockService := service.NewMockCustomerService(ctrl)
	ch = CustomerHandler{service: mockService}
	router = mux.NewRouter()
	return func() {
		defer ctrl.Finish()
	}
}

func Test_should_return_customers_with_status_code_200(t *testing.T) {
	//Arrange
	setup(t)
	ctrl := gomock.NewController(t)
	dummyCustomers := []domain.Customer{
		{ID: "1001", Name: "Jefferson", City: "New Delhi", Zipcode: "11001", DateofBirth: "2000-01-01", Status: "1"},
		{ID: "1002", Name: "Andrea", City: "New Delhi", Zipcode: "11001", DateofBirth: "2000-01-01", Status: "1"},
	}
	defer ctrl.Finish()
	mockService := service.NewMockCustomerService(ctrl)
	mockService.EXPECT().GetAllCustomers("").Return(dummyCustomers, nil)
	ch := CustomerHandler{service: mockService}
	router := mux.NewRouter()
	router.HandleFunc("/customers", ch.getAllCustomers)
	request, _ := http.NewRequest(http.MethodGet, "/customers", nil)
	// Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	// Asert
	if recorder.Code != http.StatusOK {
		t.Error("Failed while testing the status code")
	}
}

func Test_should_return_status_500_with_error_message(t *testing.T) {
	//Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockService := service.NewMockCustomerService(ctrl)
	mockService.EXPECT().GetAllCustomers("").Return(nil, errs.NewUnexpectedError("some database error"))
	ch := CustomerHandler{service: mockService}
	router := mux.NewRouter()
	router.HandleFunc("/customers", ch.getAllCustomers)
	request, _ := http.NewRequest(http.MethodGet, "/customers", nil)
	// Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	// Asert
	if recorder.Code != http.StatusInternalServerError {
		t.Error("Failed while testing the status code")
	}
}
