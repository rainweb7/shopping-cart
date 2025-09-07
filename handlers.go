package handlers

import (
	"fmt"
	"net/http"
)

// Example handler
func GetProducts(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, `{"products": ["Laptop", "Phone", "Shoes"]}`)
}
