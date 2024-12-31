package handlers

import (
	"ecomerce/data"
	"ecomerce/db"
	"ecomerce/utils"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func SigninHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		RenderPage(w, r, data.PageData{Title: "Signin", Data: nil, User: data.User{}})
		return
	case http.MethodPost:
		singInUser(w, r)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

func singInUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	r.ParseForm()

	email := r.FormValue("email")
	password := r.FormValue("password")

	if !CheckInputValues([]string{email, password}) {
		err := data.ErrorResponse{
			Code:      http.StatusBadRequest,
			Error:     "Bad Request",
			Msg:       "All Fields Must Be Filled",
			Redirect:  "/signin",
			Directive: "Try Again",
		}
		RenderPage(w, r, data.PageData{Title: "Error", Data: err, User: data.User{}})
		return
	}

	user, err := db.GetUser(email)

	if err != nil {
		log.Println("User Not Found:", err)
		errResponse := data.ErrorResponse{
			Code:      http.StatusUnauthorized,
			Error:     "Unauthorized",
			Msg:       "User Not Found",
			Redirect:  "/signup",
			Directive: "Create Account?",
		}
		RenderPage(w, r, data.PageData{Title: "Error", Data: errResponse, User: data.User{}})
		return
	}

	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		log.Println("Incorrect password")
		errResponse := data.ErrorResponse{
			Code:      http.StatusUnauthorized,
			Error:     "Unauthorized",
			Msg:       "Invalid email or password",
			Redirect:  "/signin",
			Directive: "Try Again",
		}
		RenderPage(w, r, data.PageData{Title: "Error", Data: errResponse, User: user})
		return
	}

	utils.CreateUserSession(w, r, user)
	fmt.Println(user.Role)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
