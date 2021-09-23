package app

import (
	"Banking/domain"
	"Banking/service"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
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

	ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	// define routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	// starting server
	address := os.Getenv("server_address")
	port := os.Getenv("server_port")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))

}
