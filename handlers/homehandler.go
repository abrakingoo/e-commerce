package handlers

import (
	"ecomerce/data"
	"ecomerce/db"
	"ecomerce/utils"
	"log"
	"net/http"
	"sync"
)

var (
	Products     []data.Product
	ProductsLock sync.RWMutex // Mutex for synchronizing access to Products
)

func LoadProductsCache() error {
	ProductsLock.Lock()
	defer ProductsLock.Unlock()

	// Fetch products from the database
	products, err := db.FetchProducts()
	if err != nil {
		return err
	}

	Products = products
	return nil
}

func GetCachedProducts() []data.Product {
	ProductsLock.RLock()
	defer ProductsLock.RUnlock()

	// Return a copy of the cached products to avoid accidental modification
	cachedProducts := make([]data.Product, len(Products))
	copy(cachedProducts, Products)
	return ReverseSlice(cachedProducts)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve user from session
	user, loggedIn := utils.GetUserFromSession(r)
	if !loggedIn {
		user = data.User{}
	}

	// Fetch products from cache
	cachedProducts := GetCachedProducts()

	// If cache is empty, reload from database
	if len(cachedProducts) == 0 {
		if err := LoadProductsCache(); err != nil {
			log.Printf("Error loading products cache: %v", err)
			http.Error(w, "Unable to load products. Please try again later.", http.StatusInternalServerError)
			return
		}
		cachedProducts = GetCachedProducts()
	}

	// Render the page
	dataToRender := data.PageData{
		Title: "Home",
		Data:  cachedProducts,
		User:  user,
	}

	if len(cachedProducts) == 0 {
		log.Println("No products found")
		dataToRender.Data = nil
	}

	if err := RenderPage(w, r, dataToRender); err != nil {
		log.Printf("Error rendering page: %v", err)
		http.Error(w, "An error occurred while rendering the page.", http.StatusInternalServerError)
	}
}

func ReverseSlice(slice []data.Product) []data.Product {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}