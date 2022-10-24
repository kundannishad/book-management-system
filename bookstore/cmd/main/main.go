package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kundannishad/bookstore/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	log.Fatal(http.ListenAndServe(":8000", r))
}
