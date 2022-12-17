package app

import (
	"encoding/json"
	"github.com/SupratickDey/banking/service"
	"github.com/gorilla/mux"
	"net/http"
)

// CustomerHandlers - dependency to customer handler
type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	customers, err := ch.service.GetAllCustomers(status)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, customers)
	}
}

func (ch *CustomerHandlers) getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]
	customer, err := ch.service.GetCustomer(id)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, customer)
	}
}

func writeResponse(w http.ResponseWriter, code int, data any) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		panic(err)
	}
}
