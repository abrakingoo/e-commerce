package handlers

import (
	"ecomerce/data"
	"ecomerce/utils"
	"log"
	"net/http"
	"path/filepath"
)

func AdminHandler(w http.ResponseWriter, r *http.Request) {

	if filepath.Ext(r.URL.Path) == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	
	user, ok := utils.GetUserFromSession(r)
	if !ok {
		user = data.User{}
	}
	if err := RenderPage(w, r, data.PageData{Title: "admin", Data: nil, User: user}); err != nil {
		log.Println(err)
	}
}
