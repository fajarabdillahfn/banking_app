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

func Start() {
	router := mux.NewRouter()

	// wiring
	ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryStub())}

	// define routes
	router.HandleFunc("/customers", ch.GetAllCustomers).Methods(http.MethodGet)

	// starting server
	fmt.Println("Listening on Port", port)
	log.Fatal(http.ListenAndServe(port, router))
}
