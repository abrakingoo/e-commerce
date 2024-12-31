package main

import (
	"ecomerce/db"
	"ecomerce/handlers"
	"log"
	"net/http"
	"path/filepath"
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
	// Handle static assets
	http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
		// Block directory browsing
		if filepath.Ext(r.URL.Path) == "" {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		// Serve the requested file from the static directory
		http.ServeFile(w, r, "."+r.URL.Path)
	})

	// Other routes
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
	http.HandleFunc("/checkout", handlers.CheckoutHandler)
	http.HandleFunc("/admin", handlers.AdminHandler)
	http.HandleFunc("/upload", handlers.UploadHandler)
	http.HandleFunc("/stock", handlers.StockHandler)

	log.Println("Server running on https://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
