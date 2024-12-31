package handlers

import (
	"ecomerce/db"
	"log"
	"net/http"
)

func UpdateItemHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    id := r.FormValue("id")
    title := r.FormValue("title")
    price := r.FormValue("price")
    category := r.FormValue("category")
    total := r.FormValue("total")
    description := r.FormValue("description")

    // Update the database with the new values
    err := db.UpdateProduct(id, title, price, category, total, description)
    if err != nil {
        log.Println("Error updating product:", err)
        http.Error(w, "Unable to update product", http.StatusInternalServerError)
        return
    }

    // Redirect back to the stock page
    http.Redirect(w, r, "/stock", http.StatusSeeOther)
}
