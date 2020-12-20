package controllers

import (
	"bytes"
	"devbookapp/src/responses"
	"encoding/json"
	"fmt"
	"net/http"
)

// CreateUser calls API to create user
func CreateUser(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	user, error := json.Marshal(map[string]string{
		"name": r.FormValue("name"),
		"nick": r.FormValue("nick"),
		"email": r.FormValue("email"),
		"password": r.FormValue("password"),
	})
	if error != nil {
		responses.JsonResponse(w, http.StatusBadRequest, responses.ResponseError{Error: error.Error()})
		return
	}

	response, error := http.Post(
		"http://localhost:5000/users",
		"application/json",
		bytes.NewBuffer(user))
	if error != nil {
		responses.JsonResponse(w, http.StatusInternalServerError, responses.ResponseError{Error: "Falha na API"})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		fmt.Println("Resposta com erro")
		responses.ErrorResponseResolver(w, response)
	} else {
		fmt.Printf("\nResposta com sucesso. Código %d", response.StatusCode)
		responses.JsonResponse(w, 204, nil) // fixed 204 for jquery not require a json body as response
	}
}