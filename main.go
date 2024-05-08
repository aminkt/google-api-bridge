package main

import (
	"github.com/aminkt/google-api-bridge/controllers"
	"net/http"
)

func main() {
	http.HandleFunc("/verify-token", controllers.VerifyOathTokenAction)
	http.ListenAndServe(":8080", nil)
}
