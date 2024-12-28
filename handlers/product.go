package handlers

import (
	"ecomerce/data"
	"ecomerce/utils"
	"net/http"
	"strconv"
)

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	user, ok := utils.GetUserFromSession(r)
	if !ok  {
		user = data.User{}
	}
	id, _ := strconv.Atoi((r.URL.Query().Get("id")))
	for _, product := range Products {
		if id == product.Id {
			RenderPage(w, data.PageData{Title: "product", Data: product, User: user})
			return
		}
	}

	RenderPage(w, data.PageData{Title: "product", Data: nil, User: user})
}