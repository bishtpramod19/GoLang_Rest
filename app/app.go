package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {

	// mux := http.NewServeMux()
	mux := mux.NewRouter

	mux.HandleFunc("/greet", greet)
	mux.HandleFunc("/customers", getAllCustomers)

	log.Fatal(http.ListenAndServe("localhost:8000", mux))

}
