package main

import (
	"fmt"
	"net/http"

	"github.com/asadbek21coder/http-server/handlers"
)

// PORT
const PORT = "8080"

func main() {

	// Endpoints

	// Home page endpoints
	http.HandleFunc("/", handlers.GetHomePage)

	// Users endpoints
	http.HandleFunc("/users", handlers.HandleUsers)

	// Server
	fmt.Println("Server is started working on http://localhots:" + PORT)
	http.ListenAndServe(":"+PORT, nil)
}
