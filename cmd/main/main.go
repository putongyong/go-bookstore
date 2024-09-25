package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/putongyong/go-bookstore/pkg/config"
	"github.com/putongyong/go-bookstore/pkg/models"
	"github.com/putongyong/go-bookstore/pkg/routes"
)

func main() {
	// Initialize the database connection
	config.Connect()

	// Set the db connection in the models package (this avoids cyclic imports)
	models.SetDB(config.GetDB())

	// Initialize the router
	router := mux.NewRouter()
	routes.RegisterBookstoreRoutes(router)
	http.Handle("/", router)

	// Start the server
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
