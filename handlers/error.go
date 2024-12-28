package handlers

import (
	"ecomerce/data"
	"net/http"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request) {
	RenderPage(w, data.PageData{Title: "Error", Data: "Error"})
}