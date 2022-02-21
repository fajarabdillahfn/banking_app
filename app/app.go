package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/fajarabdillahfn/banking_app/domain"
	"github.com/fajarabdillahfn/banking_app/service"
	"github.com/gorilla/mux"
)

func sanityCheck() {
	if os.Getenv("SERVER_ADDRESS") == "" ||
	os.Getenv("SERVER_PORT") == "" {
		log.Fatal("Environment variable not defined....")
	}
}

/* A function that starts the server. */
func Start() {

	sanityCheck()

	router := mux.NewRouter()

	// wiring
	// ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	// define routes
	router.HandleFunc("/customers", ch.GetAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customer/{customer_id:[0-9]+}", ch.GetCustomer).Methods(http.MethodGet)

	// starting server
	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))
}
