package middlewares

import (
	"devbookapp/src/security"
	"net/http"
)

// CheckAuthenticatedRequest checks if an authentication cookie is present in request
func CheckAuthenticatedRequest(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if _, error := security.ReadAuthCookie(r); error != nil {
			http.Redirect(w, r, "/login", 302)
			return
		}
		next(w, r)
	}
}
