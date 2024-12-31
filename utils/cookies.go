package utils

import (
	"ecomerce/data"
	// "log"
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
	session.Values["role"] = user.Role
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
		// log.Println("Error retrieving session:", err)
		return user, false
	}

	// Safely retrieve and type-assert each session value
	if firstname, ok := session.Values["firstname"].(string); ok {
		user.FName = firstname
	} else {
		// log.Println("Missing or invalid 'firstname' in session")
		return user, false
	}

	if lastname, ok := session.Values["lastname"].(string); ok {
		user.LName = lastname
	} else {
		// log.Println("Missing or invalid 'lastname' in session")
		return user, false
	}

	if email, ok := session.Values["email"].(string); ok {
		user.Email = email
	} else {
		// log.Println("Missing or invalid 'email' in session")
		return user, false
	}

	if id, ok := session.Values["id"].(string); ok {
		user.Id = id
	} else {
		// log.Println("Missing or invalid 'id' in session")
		return user, false
	}

	if phone, ok := session.Values["phonenumber"].(string); ok {
		user.Phone = phone
	} else {
		// log.Println("Missing or invalid 'phonenumber' in session")
		return user, false
	}

	if role, ok := session.Values["role"].(string); ok {
		user.Role = role
	} else {
		// log.Println("Missing or invalid 'role' in session")
		return user, false
	}

	if cart, ok := session.Values["cart"].([]string); ok {
		user.Cart = len(cart)
	} else {
		// log.Println("Missing or invalid 'cart' in session")
		return user, false
	}

	return user, true
}
