package data

import(
	
)

type Rate struct {
	Rating float64 `json:"rating"`
	Count int `json:"count"`
}

type Product struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Price float64 `json:"price"`
	Description string `json:"description"`
	Category string `json:"category"`
	Image string `json:"image"`
	Rating Rate `json:"rating"`
}