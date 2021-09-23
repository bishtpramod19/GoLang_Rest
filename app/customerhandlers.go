package app

import (
	"Banking/service"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

//dependency on primary port
type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(rw http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	customers, err := ch.service.GetAllCustomers(status)

	if err != nil {
		writeResponse(rw, err.Code, err.AsMessage())
	} else {
		writeResponse(rw, http.StatusOK, customers)
	}
}

func (ch *CustomerHandlers) getCustomer(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]

	customer, err := ch.service.GetCustomer(id)
	if err != nil {
		writeResponse(rw, err.Code, err.AsMessage())
	} else {
		writeResponse(rw, http.StatusOK, customer)
	}
}

func writeResponse(rw http.ResponseWriter, code int, data interface{}) {
	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(code)
	if err := json.NewEncoder(rw).Encode(data); err != nil {
		panic(err)

	}
}
