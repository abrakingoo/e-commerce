package main

import (
	"ecomerce/handlers"
	"log"
	"net/http"
)


func main() {

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", handlers.HomeHandler)
	http.ListenAndServe(":8080", nil)
	log.Println("Server running on https://localhost:8080")
}