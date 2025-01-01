package handlers

import (
	"ecomerce/data"
	"ecomerce/db"
	"ecomerce/utils"
	"log"
	"net/http"
)

func StockHandler(w http.ResponseWriter, r *http.Request) {

	user, ok := utils.GetUserFromSession(r)

	if !ok || user.Role != "admin" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	products, err := db.FetchProducts()
	if err != nil {
		log.Println(err)
		return
	}

	RenderPage(w, r, data.PageData{Title: "stock", Data: products, User: user})
}
