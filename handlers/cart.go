package handlers

import (
	"ecomerce/data"
	"ecomerce/utils"
	"encoding/json"

	// "fmt"
	"net/http"
	"strconv"
)

var id string

func CartHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodPost:
		id = r.URL.Query().Get("id")
		_, loggedIn := utils.GetUserFromSession(r)
		if !loggedIn {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		session, _ := utils.Store.Get(r, "session")
		cart, ok := session.Values["cart"].([]string)
		if !ok {
			cart = []string{}
		}
		cart = append(cart, id)
		session.Values["cart"] = cart
		session.Save(r, w)
		count := len(cart)
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application-json")
		var response = map[string]int{
			"Count": count,
		}

		resBody, err := json.Marshal(response)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(resBody)
		// fmt.Println(cart)

	case http.MethodGet:
		user, loggedIn := utils.GetUserFromSession(r)
		if !loggedIn {
			user = data.User{}
		}
		session, _ := utils.Store.Get(r, "session")
		cart, ok := session.Values["cart"].([]string)
		if !ok {
			cart = []string{}
		}
		// fmt.Println(cart)
		var userCart []data.Product
		if len(cart) > 0 {
			for i := len(cart) - 1; i >= 0; i-- {
				intId, _ := strconv.Atoi(cart[i])
				for _, prod := range Products {
					if intId == prod.Id {
						userCart = append(userCart, prod)
					}
				}
			}
			RenderPage(w, data.PageData{Title: "cart", Data: userCart, User: user})
		} else {

			RenderPage(w, data.PageData{Title: "cart", Data: nil, User: user})
		}

	}
}

func RemoveHandler(w http.ResponseWriter, r *http.Request) {
	user, loggedIn := utils.GetUserFromSession(r)
	if !loggedIn {
		user = data.User{}
	}

	session, _ := utils.Store.Get(r, "session")
	cart, ok := session.Values["cart"].([]string)
	if !ok {
		cart = []string{}
	}

	id := r.URL.Query().Get("id")

	for i := len(cart)-1; i >= 0; i-- {
		if cart[i] == id {
			cart = append(cart[:i], cart[i+1:]...)
			break 
		}
	}

	session.Values["cart"] = cart
	session.Save(r, w)

	RenderPage(w, data.PageData{Title: "cart", Data: cart, User: user})
}
