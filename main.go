package main

import (
	"ecomerce/db"
	"ecomerce/handlers"
	"log"
	"net/http"
)

func init() {
	if err := db.InitDB(); err != nil {
		log.Fatal(err.Error())
	}

	if err := db.CreateTables(); err != nil {
		log.Fatal(err.Error())
	}

}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/signin", handlers.SigninHandler)
	http.HandleFunc("/signout", handlers.SignOutHandler)
	http.HandleFunc("/signup", handlers.SignupHandler)
	http.HandleFunc("/error", handlers.ErrorHandler)
	http.HandleFunc("/cart", handlers.CartHandler)
	http.HandleFunc("/profile", handlers.ProfileHandler)
	http.HandleFunc("/remove", handlers.RemoveHandler)
	http.HandleFunc("/product", handlers.ProductsHandler)
	http.HandleFunc("/hotdeals", handlers.HotDealsHandler)
	log.Println("Server running on https://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
