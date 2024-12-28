package handlers

import (
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
		log.Fatalf("Failed to parse template: %v", err)
		return
	}
}

func RenderPage(w http.ResponseWriter, data interface{}) error {
	if  err = tpl.ExecuteTemplate(w, "index.html", data); err != nil {
		return err
	}
	return nil
}