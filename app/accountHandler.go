package app

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jeffleon/banking-hexarch/dto"
	"github.com/jeffleon/banking-hexarch/service"
)

type AccountHandler struct {
	service service.AccountService
}

func (a AccountHandler) NewAccount(w http.ResponseWriter, r *http.Request) {
	var request dto.NewAccountRequest
	vars := mux.Vars(r)
	customerID := vars["customer_id"]
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		request.CustomerID = customerID
		account, appError := a.service.NewAccount(request)
		if appError != nil {
			writeResponse(w, appError.Code, appError.Message)
		} else {
			writeResponse(w, http.StatusCreated, account)
		}
	}

}
