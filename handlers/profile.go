package handlers

import (
	"ecomerce/data"
	"ecomerce/utils"
	"fmt"
	"net/http"
)

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	user, ok := utils.GetUserFromSession(r)
	if !ok {
		http.Redirect(w, r, "/signin", http.StatusSeeOther)
		return
	}

	// Extract the 'name' query parameter
	name := r.URL.Query().Get("name")

	// Determine the data to pass to the template
	var page interface{}
	var res []data.Order

	switch name {
	case "account":
		page = "account"
	case "orders":
		res = utils.GetOrders(user)
		page = "orders"
	default:
		if name != "" {
			fmt.Printf("Unhandled profile section: %s\n", name)
		}
		page = nil
	}

	RenderPage(w, r, data.PageData{
		Title: "profile",
		Page:  page,
		Data:  reverseSlice(res),
		User:  user,
	})
}

func reverseSlice(slice []data.Order) []data.Order {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}
