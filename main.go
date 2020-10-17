package main

import (
	"dashboard-pi/api/handler"
	"log"
	"net/http"
)

func handleRequests() {
	http.Handle("/CPU", handler.RootHandler(handler.CPUHandler))
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	handleRequests()
}
