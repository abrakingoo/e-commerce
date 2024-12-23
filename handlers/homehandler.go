package handlers

import (
	"net/http"
	"html/template"
)

var (
	tpl *template.Template
	err error
)

func init() {
	tpl, err = template.ParseGlob("templates/*.html")
	if err != nil {
		return
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// url := "https://fakestoreapi.com/products"

	// res, err := http.Get(url)

	// if err != nil {
	// 	w.Write([]byte("Error Fetching Products"))
	// }

	// defer res.Body.Close()

	// resBody, _ := io.ReadAll(res.Body)

	// products := []data.Product{}

	// err = json.Unmarshal(resBody, &products)

	// if err != nil {
	// 	w.Write([]byte("Error Unmarshalling Products"))
	// }

	tpl.ExecuteTemplate(w, "index.html", nil)
}