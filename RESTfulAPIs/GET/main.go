package main

import (
	"encoding/json"
	"net/http"
)

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func getProductHandler(w http.ResponseWriter, r *http.Request) {
	product := Product{ID: 1, Name: "Laptop", Price: 999.99}
	json.NewEncoder(w).Encode(product)
}

func main() {
	http.HandleFunc("/product", getProductHandler)
	http.ListenAndServe(":8080", nil)
}
