package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/jeffleon/banking-hexarch/domain"
	"github.com/jeffleon/banking-hexarch/service"
	"github.com/jmoiron/sqlx"
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
	dbClient := getDbClient()
	// Repositories
	customerRepositoryDb := domain.NewCustomerRepositoryDB(dbClient)
	accountsRepositoryDb := domain.NewAccountRepositoryDb(dbClient)
	// Services
	customerService := service.NewCustomerService(customerRepositoryDb)
	accountService := service.NewAccountService(accountsRepositoryDb)
	// Handlers
	ch := CustomerHandler{service: customerService}
	ah := AccountHandler{service: accountService}
	//Routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customer/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customer/{customer_id:[0-9]+}/account", ah.NewAccount).Methods(http.MethodPost)
	//starting server
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))
}

func getDbClient() *sqlx.DB {
	dbUser := os.Getenv("DB_USER")
	dbPasswd := os.Getenv("DB_PASSWD")
	dbAddr := os.Getenv("DB_ADDR")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPasswd, dbAddr, dbPort, dbName)
	client, err := sqlx.Open("mysql", dataSource)
	// client, err := sqlx.Open("mysql", "root:secret@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}
