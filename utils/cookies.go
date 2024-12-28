package utils

import (
	"ecomerce/data"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

var Store = sessions.NewCookieStore([]byte("session"))

// CreateUserSession creates a user session and stores the necessary user data in the session.
func CreateUserSession(w http.ResponseWriter, r *http.Request, user data.User) error {
	// Create a new session or retrieve the existing one
	session, err := Store.Get(r, "session")
	if err != nil {
		http.Error(w, "Error retrieving session", http.StatusInternalServerError)
		return err
	}

	// Set session values (Don't store sensitive info like password in session)
	session.Values["firstname"] = user.FName
	session.Values["lastname"] = user.LName
	session.Values["email"] = user.Email
	session.Values["id"] = user.Id
	session.Values["phonenumber"] = user.Phone
	session.Values["cart"] = []string{}

	// Optionally, set the session expiration time
	session.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   24 * 60 * 60, // Expires in 24 hours
		HttpOnly: true,
		Secure:   true, // Ensures the cookie is only sent over HTTPS
		SameSite: http.SameSiteStrictMode,
	}

	// Save the session to the response
	if err := session.Save(r, w); err != nil {
		http.Error(w, "Error saving session", http.StatusInternalServerError)
		return err
	}

	// No need to set a separate session cookie. `gorilla/sessions` automatically does this for you.

	return nil
}

// GetUserFromSession retrieves the user from the session. It returns the user data and a boolean indicating success.
func GetUserFromSession(r *http.Request) (data.User, bool) {
	var user data.User

	// Retrieve the session
	session, err := Store.Get(r, "session")
	if err != nil {
		log.Println("Error retrieving session:", err)
		return user, false
	}

	// Check if user data exists in the session
	firstname, ok := session.Values["firstname"].(string)
	if !ok {
		log.Println("No user found in session")
		return user, false
	}

	// Populate the user struct with session values
	user.FName = firstname
	user.LName = session.Values["lastname"].(string)
	user.Email = session.Values["email"].(string)
	user.Id = session.Values["id"].(string)
	user.Phone = session.Values["phonenumber"].(string)
	user.Cart = len(session.Values["cart"].([]string))

	return user, true
}