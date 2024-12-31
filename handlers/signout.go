package handlers

import (
	"ecomerce/utils"
	"net/http"
	"time"
)

func SignOutHandler(w http.ResponseWriter, r *http.Request) {
	// Get the session
	session, _ := utils.Store.Get(r, "session")

	// Clear session values
	session.Values["firstname"] = nil
	session.Values["lastname"] = nil
	session.Values["email"] = nil
	session.Values["id"] = nil
	session.Values["phonenumber"] = nil
	session.Values["role"] = nil
	session.Values["cart"] = nil

	// Expire the session cookie by setting MaxAge to -1
	http.SetCookie(w, &http.Cookie{
		Name:    "session",
		Value:   "",
		MaxAge:  -1, // Expire the cookie immediately
		Expires: time.Now().Add(-time.Hour),
		Secure:  true, // Secure flag for HTTPS
		HttpOnly: true, // Ensure the cookie is not accessible via JavaScript
		SameSite: http.SameSiteStrictMode, // Strict SameSite policy
	})

	// Save the session after clearing it
	session.Save(r, w)

	// Redirect to the homepage or login page
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
