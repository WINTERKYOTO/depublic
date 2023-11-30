package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/depublic/depublic/internal/config"
	"github.com/depublic/depublic/internal/http/handler"
	"github.com/depublic/depublic/repository"
	"github.com/depublic/depublic/service"
	"github.com/gorilla/mux"
)

func main() {
	// Load the configuration
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Initialize the database
	db, err := repository.NewDatabase(config.Database)
	if err != nil {
		log.Fatal(err)
	}

	// Create the repositories
	productRepository := repository.NewProductRepository(db)

	// Create the services
	productService := service.NewProductService(productRepository)

	// Create the handlers
	productHandler := handler.NewProductHandler(productService)

	// Create the router
	router := mux.NewRouter()

	// Register the routes
	router.HandleFunc("/products", productHandler.GetProducts).Methods("GET")
	router.HandleFunc("/products/{id}", productHandler.GetProduct).Methods("GET")
	router.HandleFunc("/products", productHandler.CreateProduct).Methods("POST")
	router.HandleFunc("/products/{id}", productHandler.UpdateProduct).Methods("PUT")
	router.HandleFunc("/products/{id}", productHandler.DeleteProduct).Methods("DELETE")

	// Start the server
	log.Println("Starting server...")
	http.ListenAndServe(fmt.Sprintf(":%s", config.Port), router)
}
