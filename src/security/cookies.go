package security

import (
	"devbookapp/src/config"
	"github.com/gorilla/securecookie"
	"net/http"
	"time"
)

var s *securecookie.SecureCookie

// ConfigureSecureCookie configures a secure cookie to keep auth token
func ConfigureSecureCookie() {
	s = securecookie.New(config.HashKey, config.BlockKey)
}

func SetAuthCookie(w http.ResponseWriter, token, userId string) error {
	data := map[string]string {
		"id" : userId,
		"token": token,
	}

	encodedData, error := s.Encode("SID", data)
	if error != nil {
		return error
	}

	http.SetCookie(w, &http.Cookie{
		Name: "SID",
		Value: encodedData,
		Path: "/",
		HttpOnly: true,
	})

	return nil
}

// ReadAuthCookie returns values stored in auth cookie
func ReadAuthCookie(r *http.Request) (map[string]string, error) {
	cookie, error := r.Cookie("SID")
	if error != nil {
		return nil, error
	}

	cookieValues := make(map[string]string)
	if error = s.Decode("SID", cookie.Value, &cookieValues); error != nil {
		return nil, error
	}

	return cookieValues, nil
}

// DeleteCookieValue removes authentication cookie value
func DeleteCookieValue(w http.ResponseWriter){
	http.SetCookie(w, &http.Cookie{
		Name: "SID",
		Value: "",
		Path: "/",
		HttpOnly: true,
		Expires: time.Unix(0, 0),
	})
}
