package server

import (
	"encoding/json"
	"go-customer-ms/pkg/model"

	"net/http"

	"github.com/gorilla/mux"
)

type api struct {
	router http.Handler
}

// Server ...
type Server interface {
	Router() http.Handler
}

// New ...
func New() Server {
	a := &api{}

	r := mux.NewRouter()
	//Configurando las rutas que debe resolver
	r.HandleFunc("/customers", a.fetchAllCustomers).Methods(http.MethodGet)
	// Note como Gorilla mux permite colocar expresione regulares para establecer reglas en los parámetros que
	// se pasen
	r.HandleFunc("/customers/{ID:[a-zA-Z0-9_]+}", a.fetchCustomer).Methods(http.MethodGet)
	a.router = r
	return a
}

func (a *api) Router() http.Handler {
	return a.router
}

// Customers ...
type Customers []model.Customer

func (a *api) fetchAllCustomers(w http.ResponseWriter, r *http.Request) {
	/*
		customers, _ := a.repository.FetchGophers()

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)*/
	customers := Customers{
		model.Customer{ID: "c123", Name: "Juli", Age: 32},
	}
	json.NewEncoder(w).Encode(customers)
}

func (a *api) fetchCustomer(w http.ResponseWriter, r *http.Request) {
	/*vars := mux.Vars(r)
	gopher, err := a.repository.FetchGopherByID(vars["ID"])
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusNotFound) // We use not found for simplicity
		json.NewEncoder(w).Encode("Gopher Not found")
		return
	}*/

	json.NewEncoder(w).Encode(model.Customer{ID: "c111523", Name: "Juliano", Age: 32})
}
