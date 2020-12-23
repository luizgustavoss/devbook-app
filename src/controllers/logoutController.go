package controllers

import (
	"devbookapp/src/security"
	"net/http"
)

// Logout logs user out
func Logout(w http.ResponseWriter, r *http.Request) {
	security.DeleteCookieValue(w)
	http.Redirect(w, r, "/login", 302)
}