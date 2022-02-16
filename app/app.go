package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fajarabdillahfn/banking_app/domain"
	"github.com/fajarabdillahfn/banking_app/service"
	"github.com/gorilla/mux"
)

const port = ":8000"

/* A function that starts the server. */
func Start() {
	router := mux.NewRouter()

	// wiring
	// ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	// define routes
	router.HandleFunc("/customers", ch.GetAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customer/{customer_id:[0-9]+}", ch.GetCustomer).Methods(http.MethodGet)

	// starting server
	fmt.Println("Listening on Port", port)
	log.Fatal(http.ListenAndServe(port, router))
}
