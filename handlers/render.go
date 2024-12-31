package handlers

import (
	"ecomerce/data"
	"ecomerce/utils"
	"html/template"
	"log"
	"net/http"
)

var (
	tpl *template.Template
	err error
)

func init() {
	tpl, err = template.ParseGlob("templates/*.html")
	if err != nil {
		log.Panicf("Failed to parse template: %v", err)
	}
}

func RenderPage(w http.ResponseWriter, r *http.Request, input interface{}) error {

	user, ok := utils.GetUserFromSession(r)

	if !ok {
		user = data.User{}
	}


	if user.Role == "admin" {
		if  err = tpl.ExecuteTemplate(w, "admin.html", input); err != nil {
			return err
		}
		return nil
	}

	if  err = tpl.ExecuteTemplate(w, "index.html", input); err != nil {
		return err
	}
	return nil
}