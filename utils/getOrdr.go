package utils

import (
	"database/sql"
	"ecomerce/data"
	"ecomerce/db"
	"fmt"
	"log"
	"strings"
)

func GetOrders(user data.User) []data.Order {
	var Products = []data.Product{}
	var err error

	Products, err = db.FetchProducts()
	if err != nil {
		log.Println(err)
		return nil
	}
	var orders []data.Order

	stm := `SELECT id, userid, productid, total, status FROM orders WHERE userid = ?`

	rows, err := db.DB.Query(stm, user.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No Orders Found")
			return nil
		}
		fmt.Println("Error querrying db", err)
		return nil
	}

	defer rows.Close() // Ensure rows are closed

	// Iterate through the result set
	for rows.Next() {
		var order data.Order
		var productid string
		err := rows.Scan(&order.Id, &order.UserId, &productid, &order.Total, &order.Status)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			continue
		}
		ids := strings.Split(strings.TrimSpace(productid), ",")
		for _, id := range ids {
			for _, prod := range Products {
				if prod.Id == id {
					order.Products = append(order.Products, prod)
				}
			}
		}
		orders = append(orders, order)
	}

	// Check for any errors encountered during iteration
	if err := rows.Err(); err != nil {
		fmt.Println("Error iterating rows:", err)
		return nil
	}
	return orders
}
