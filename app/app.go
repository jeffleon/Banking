package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jeffleon/banking-hexarch/domain"
	"github.com/jeffleon/banking-hexarch/service"
)

func sanityCheck() {
	if os.Getenv("SERVER_PORT") == "" ||
		os.Getenv("SERVER_ADDRESS") == "" {
		log.Fatal("Environment variable not defined ...")
	}
}

func Start() {
	sanityCheck()
	router := mux.NewRouter()
	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	//define routes
	ch := CustomerHandler{service: service.NewCustomerService(domain.NewCustomerRepositoryDB())}
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customer/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)
	//starting server
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))
}
