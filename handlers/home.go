package handlers

import (
	"fmt"
	"net/http"
)

// Handlers
func GetHomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to Home page of our website")
}
