package handlers

import (
	// "ecomerce/data"
	"ecomerce/db"
	"ecomerce/utils"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gofrs/uuid"
)

func CheckoutHandler(w http.ResponseWriter, r *http.Request) {
	user, ok := utils.GetUserFromSession(r)

	if !ok {
		http.Redirect(w, r, "/signin", http.StatusSeeOther)
		return
	}

	session, err := utils.Store.Get(r, "session")

	if err != nil {
		fmt.Println(err)
		return
	}

	cart, ok := session.Values["cart"].([]string)

	if !ok {
		cart = []string{}
	}

	if len(cart) != 0 {

	var total float64

	Products := GetCachedProducts()
	for _, id := range cart {
		for _, p := range Products {
			if id == p.Id {
				price, _ := strconv.Atoi(p.Price)
				total += float64(price)
			}
		}
	}
	orderId , _ := uuid.NewV4()
	id := orderId.String()[:8]
	products := strings.Join(cart, ",")
		stm := `INSERT INTO orders(id, userid, productid, total) VALUES(?,?,?,?)`
		_, err = db.DB.Exec(stm, id, user.Id, products, total)
	
		if err != nil {
			fmt.Println(err)
		}
	}

	session.Values["cart"] = []string{}
	session.Save(r, w);
	
	// RenderPage(w, data.PageData{Title: "Home", Data: Products, User: user})
	http.Redirect(w, r, "/profile?name=orders", http.StatusSeeOther)
	// fmt.Println("Your Cart: ", strings.Join(cart, ","))
}
