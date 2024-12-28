package handlers

import (
	"ecomerce/data"
	"ecomerce/utils"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

var Products = []data.Product{}

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	user, loggedIn := utils.GetUserFromSession(r)
	if !loggedIn {
		user = data.User{}
	}

	file, err := os.ReadFile("db/products.json")

	if err != nil {
		log.Println(err.Error())
	}

	if err = json.Unmarshal(file, &Products); err != nil {
		log.Println(err.Error())
	}

	// If products are empty, handle that case by rendering a message
	if len(Products) == 0 {
		log.Println("No products found")
		RenderPage(w, data.PageData{Title: "Home", Data: "No products available at the moment", User: user})
		return
	}

	// Render the page with the fetched products and user data
	RenderPage(w, data.PageData{Title: "Home", Data: Products, User: user})
}
