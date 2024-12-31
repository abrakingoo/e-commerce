package handlers

import (
	"ecomerce/data"
	"ecomerce/utils"
	"net/http"
)

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	user, ok := utils.GetUserFromSession(r)
	if !ok {
		user = data.User{}
	}
	id := r.URL.Query().Get("id")
	for _, product := range Products {
		if id == product.Id {
			RenderPage(w, r, data.PageData{Title: "product", Data: product, User: user})
			return
		}
	}

	RenderPage(w, r, data.PageData{Title: "product", Data: nil, User: user})
}
