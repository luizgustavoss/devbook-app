package requests

import (
	"devbookapp/src/security"
	"io"
	"net/http"
)

// RequestAuthenticatedEndpoint adds a jwt token to request API endpoints
func RequestAuthenticatedEndpoint(r *http.Request, method, url string, data io.Reader) (*http.Response, error) {

	request, error := http.NewRequest(method, url, data)
	if error != nil {
		return nil, error
	}

	cookie, _ := security.ReadAuthCookie(r)
	request.Header.Add("Authorization", "Bearer " + cookie["token"])

	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		return nil, error
	}

	return response, nil
}
