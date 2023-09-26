package api

import (
	"github.com/gorilla/mux"
	"github.com/ivkoandrv/simple-rest-go/handlers"
)

func SetupRoutes(router *mux.Router) {
	// searchProducts
	router.HandleFunc("/api/products", handlers.SearchProducts).Methods("GET", "OPTIONS")
	// getProducts
	router.HandleFunc("/api/products", handlers.CreateProduct).Methods("POST", "OPTIONS")
}
