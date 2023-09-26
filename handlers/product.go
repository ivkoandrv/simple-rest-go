package handlers

import (
	"encoding/json"
	"github.com/ivkoandrv/simple-rest-go/models"
	"net/http"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newItem models.Product
	json.NewDecoder(r.Body).Decode(&newItem)
	models.Products = append(models.Products, newItem) // Use models.Items
	json.NewEncoder(w).Encode(newItem)
}

func SearchProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.Products) // Use models.Items
}
