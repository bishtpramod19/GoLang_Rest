package app

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Customer struct {
	Name    string `json:"full_name"`
	City    string `json:"city"`
	Zipcode string `json:"zipcode"`
}

func greet(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "Hi Pramod !! welcome back !")
}

func getAllCustomers(rw http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{Name: "Pramod Singh", City: "New Delhi", Zipcode: "201303"},
		{Name: "Priyanka Rawat", City: "Noida", Zipcode: "110092"},
	}

	rw.Header().Add("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(customers)

}
