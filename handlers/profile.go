package handlers

import (
	"ecomerce/data"
	"ecomerce/utils"
	"net/http"
)

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	user, ok := utils.GetUserFromSession(r)

	if !ok {
		http.Redirect(w, r, "/signin", http.StatusSeeOther)
	} else {

		RenderPage(w, data.PageData{Title: "profile", Data: nil, User: user})
	}

}