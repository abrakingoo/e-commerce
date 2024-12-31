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

	if !ok {
		user = data.User{}
	}
	
	products, err := db.FetchProducts()
	if err != nil {
		log.Println(err)
		return
	}

	RenderPage(w, r, data.PageData{Title: "stock", Data: products, User:user })
}