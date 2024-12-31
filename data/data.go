package data

type PageData struct {
	Title string
	Page  interface{}
	Data  interface{}
	User  interface{}
}

type Rate struct {
	Rating float64 `json:"rating"`
	Count  int     `json:"count"`
}

type Product struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Price       string `json:"price"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Image       string `json:"image"`
	Total       string    `json:"total"`
	Rating      int    `json:"rating"`
}

type ErrorResponse struct {
	Code      int64  `json:"code"`
	Error     string `json:"error"`
	Msg       string `json:"msg"`
	Redirect  string `json:"redirect"`
	Directive string `json:"directive"`
}

type User struct {
	Id       string
	FName    string
	LName    string
	Phone    string
	Email    string
	Password string
	Role     string
	Cart     int
}

type Order struct {
	Id       string    `json:"id"`
	UserId   string    `json:"userid"`
	Products []Product `json:"productid"`
	Total    string    `json:"total"`
	Status   bool      `json:"status"`
}
