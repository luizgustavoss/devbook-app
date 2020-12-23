package controllers

import (
	"bytes"
	"devbookapp/src/config"
	"devbookapp/src/models"
	"devbookapp/src/responses"
	"devbookapp/src/security"
	"encoding/json"
	"fmt"
	"net/http"
)

// Login logs in a user
func Login(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	credentials, error := json.Marshal(map[string]string{
		"email": r.FormValue("email"),
		"password": r.FormValue("password"),
	})
	if error != nil {
		responses.JsonResponse(w, http.StatusBadRequest, responses.ResponseError{Error: error.Error()})
		return
	}

	url := fmt.Sprintf("%s/login", config.ApiUrl)
	response, error := http.Post(
		url, "application/json",
		bytes.NewBuffer(credentials))
	if error != nil {
		responses.JsonResponse(w, http.StatusInternalServerError, responses.ResponseError{Error: error.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		fmt.Println("Resposta com erro")
		responses.ErrorResponseResolver(w, response)
	} else {
		var token models.Token
		if error = json.NewDecoder(response.Body).Decode(&token); error != nil {
			responses.JsonResponse(w, http.StatusBadRequest, responses.ResponseError{Error: error.Error()})
			return
		}

		if error = security.SetAuthCookie(w, token.Token, token.UserId); error != nil {
			responses.JsonResponse(w, http.StatusBadRequest, responses.ResponseError{Error: error.Error()})
			return
		}

		responses.JsonResponse(w, http.StatusNoContent, nil)
	}
}
