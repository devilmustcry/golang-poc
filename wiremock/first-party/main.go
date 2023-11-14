package main

import (
	"golang-poc/wiremock/first-party/controller"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", controller.HandleHello)
	http.ListenAndServe(":8081", nil)
}
