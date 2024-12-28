package data

import(
	
)

type PageData struct {
	Title string
	Data interface{}
	User interface{}
}

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

type ErrorResponse struct {
    Code  int64  `json:"code"`
    Error string `json:"error"`
    Msg   string `json:"msg"`
	Redirect string `json:"redirect"`
	Directive string `json:"directive"`
}

type User struct {
	Id string
	FName string
	LName string
	Phone string
	Email string
	Password string
	Cart int
}