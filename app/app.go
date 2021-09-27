package app

import (
	"Banking/domain"
	"Banking/service"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func sanityCheck() {
	if os.Getenv("server_address") == "" || os.Getenv("server_port") == "server_port" {
		log.Fatal("Required environment variable not defined...")
	}
}
func Start() {

	sanityCheck()

	// mux := http.NewServeMux()
	router := mux.NewRouter()

	// wiring
	//ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryStub())}

	dbClient := getDbClient()
	customerRepositoryDB := domain.NewCustomerRepositoryDb(dbClient)
	accountRepositoryDb := domain.NewAccountRepositoryDB(dbClient)

	//handlers
	ch := CustomerHandlers{service: service.NewCustomerService(customerRepositoryDB)}
	ah := AccountHandler{service: service.NewAccountService(accountRepositoryDb)}

	// define routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.NewAccount).Methods(http.MethodPost)

	router.HandleFunc("/customers/{customer_id:[0-9]+}/account/{account_id:[0-9]+}", ah.MakeTransaction).Methods(http.MethodPost)

	// starting server
	address := os.Getenv("server_address")
	port := os.Getenv("server_port")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))

}

func getDbClient() *sqlx.DB {
	dbUser := os.Getenv("db_user")
	dbPasswd := os.Getenv("db_password")
	//dbAddr := os.Getenv("db_add")
	//dbPort := os.Getenv("db_port ")
	dbName := os.Getenv("db_name")

	dataSource := fmt.Sprintf("%s:%s@/%s", dbUser, dbPasswd, dbName)
	client, err := sqlx.Open("mysql", dataSource)
	//client, err := sqlx.Open("mysql", "banking:banking@123@/banking")

	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return client
}
