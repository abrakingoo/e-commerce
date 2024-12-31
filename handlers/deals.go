package handlers

import (
	"ecomerce/data"
	"ecomerce/utils"
	"net/http"
)

func HotDealsHandler(w http.ResponseWriter, r *http.Request) {
	user, ok := utils.GetUserFromSession(r)

	if !ok {
		user = data.User{}
	}

	RenderPage(w, r, data.PageData{Title: "hotdeals", Data: nil, User: user})
}
