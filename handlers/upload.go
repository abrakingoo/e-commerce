package handlers

import (
	"ecomerce/db"
	"ecomerce/utils"
	"log"
	"net/http"
	"strings"
)

func CheckInput(arr []string) bool {
	for _, str := range arr {
		if strings.TrimSpace(str) == "" {
			return false
		}
	}
	return true
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	_, ok := utils.GetUserFromSession(r)

	if !ok {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if r.Method == http.MethodPost {

		if err := r.ParseForm(); err != nil {
			log.Println(err)
			return
		}

		name := r.FormValue("product_name")
		description := r.FormValue("product_description")
		category := r.FormValue("product_category")
		price := r.FormValue("product_price")
		image := r.FormValue("product_image")
		total := r.FormValue("product_quantity")

		if !CheckInput([]string{description, category, name, price, image, total}) {
			log.Println("All fields must not be empty")
			return
		}

		if err = db.AddProduct(name, description, category, price, image, total); err != nil {
			log.Println(err)
			return
		}

		// w.Write([]byte(name + ", " + description + ", " + category + ", " + price + ", " + image))
		LoadProductsCache()
		http.Redirect(w, r, "/upload", http.StatusSeeOther)
		return

	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
