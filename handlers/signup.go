package handlers

import (
	"ecomerce/data"
	"ecomerce/db"
	"log"
	"net/http"
	"strings"

	"github.com/gofrs/uuid"
	"github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		RenderPage(w, r, data.PageData{Title: "Signup", Data: nil, User: data.User{}})
	case http.MethodPost:
		registerNewUser(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		RenderPage(w, r, data.PageData{Title: "Error", Data: "Method Not Allowed", User: data.User{}})
	}
}

func registerNewUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fname := r.FormValue("firstname")
	lname := r.FormValue("lastname")
	number := r.FormValue("number")
	email := r.FormValue("email")
	password := r.FormValue("password")

	if !CheckInputValues([]string{fname, lname, number, email, password}) {
		err := data.ErrorResponse{
			Code:      http.StatusBadRequest,
			Error:     "Bad Request",
			Msg:       "All Fields Must Be Filled",
			Redirect:  "/signup",
			Directive: "Try Again",
		}
		RenderPage(w, r, data.PageData{Title: "Error", Data: err, User: data.User{}})
		return
	}

	hashedpassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		log.Fatal(err)
	}

	uid, err := uuid.NewV4()
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(uid, string(hashedpassword))
	er := db.AddUser(uid.String(), fname, lname, number, email, string(hashedpassword))
	if er != nil {
		if sqlerr, ok := er.(sqlite3.Error); ok && sqlerr.ExtendedCode == sqlite3.ErrConstraintUnique {
			errResp := data.ErrorResponse{
				Code:      http.StatusBadRequest,
				Error:     "Bad Request",
				Msg:       "The email alredy registered to another user",
				Redirect:  "/signup",
				Directive: "Try Again",
			}
			RenderPage(w, r, data.PageData{Title: "Error", Data: errResp, User: data.User{}})
			return
		}
	}

	http.Redirect(w, r, "/signin", http.StatusSeeOther)
}

func CheckInputValues(arr []string) bool {
	for _, str := range arr {
		if strings.TrimSpace(str) == "" {
			return false
		}
	}
	return true
}
