package app

import (
	"github.com/SupratickDey/banking/domain"
	"github.com/SupratickDey/banking/service"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"time"
)

func Start() {

	router := mux.NewRouter()

	//wiring of adapters and ports with handlers

	//ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	dbClient := getDbClient()
	ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryDb(dbClient))}
	ah := AccountHandler{service: service.NewAccountService(domain.NewAccountRepositoryDb(dbClient))}

	// define routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.NewAccount()).Methods(http.MethodPost)

	router.HandleFunc("/customers/{customer_id:[0-9]+}/account/{account_id:[0-9]+}", ah.NewAccount()).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe("localhost:8080", router))

}

func getDbClient() *sqlx.DB {
	client, err := sqlx.Open("mysql", "freedb_supratick:ENee9v9Wp$9?WZq@tcp(sql.freedb.tech:3306)/freedb_banking")
	//"freedb_supratick:ENee9v9Wp$9?WZq@tcp(sql.freedb.tech:3306)/freedb_banking"
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}
