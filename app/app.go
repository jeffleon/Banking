package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jeffleon/banking-hexarch/domain"
	"github.com/jeffleon/banking-hexarch/service"
)

func Start() {
	//mux := http.NewServeMux()
	router := mux.NewRouter()
	//define routes
	ch := CustomerHandler{service: service.NewCustomerService(domain.NewCustomerRepositoryDB())}
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customer/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)
	//starting server
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
