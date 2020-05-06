package main

// Main or entry point for our application
import (
	"log"

	"net/http"

	"go-customer-ms/pkg/server" // confuguranto gorilla mux
)

/*
func main() {
	// Si es cierto :) , con esto ya levantamos un server
	log.Fatal(http.ListenAndServe(":8080", nil))
}*/

func main() {
	s := server.New()
	log.Fatal(http.ListenAndServe(":8080", s.Router()))
}
