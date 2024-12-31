package db

import (
	"ecomerce/data"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

func Update() {
	file, err := os.Open("db/products.json")
	if err != nil {
		log.Println(err)
		return
	}

	filedata, err := io.ReadAll(file)
	if err != nil {
		log.Println(err)
		return
	}

	products := []data.Product{}

	if err = json.Unmarshal(filedata, &products); err != nil {
		log.Println(err)
		return
	}

	fmt.Println(products)
}
// FetchProducts fetches all products from the database.
func FetchProducts() ([]data.Product, error) {
	var products []data.Product
	// SQL statement to fetch products
	stm := `SELECT id, title, price, description, category, image, total, rating FROM products`

	// Execute the query
	rows, err := DB.Query(stm)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	// Iterate through rows and populate the products slice
	for rows.Next() {
		var product data.Product
		if err := rows.Scan(
			&product.Id,
			&product.Title,
			&product.Price,
			&product.Description,
			&product.Category,
			&product.Image,
			&product.Total,
			&product.Rating,
		); err != nil {
			log.Printf("error scanning row: %v", err)
			continue
		}
		products = append(products, product)
	}

	// Check for errors that may have occurred during iteration
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %v", err)
	}

	return products, nil
}
