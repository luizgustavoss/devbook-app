package controllers

import (
	"bytes"
	"devbookapp/src/config"
	"devbookapp/src/requests"
	"devbookapp/src/responses"
	"devbookapp/src/security"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
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

	url := fmt.Sprintf("%s/users", config.ApiUrl)
	response, error := http.Post(
		url, "application/json",
		bytes.NewBuffer(user))
	if error != nil {
		responses.JsonResponse(w, http.StatusInternalServerError, responses.ResponseError{Error: error.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.ErrorResponseResolver(w, response)
		return
	}
	responses.JsonResponse(w, 204, nil) // fixed 204 for jquery not require a json body as response
}

// FollowUser seguir um usuário
func FollowUser(w http.ResponseWriter, r *http.Request) {

	pathParameters := mux.Vars(r)
	userId, error := strconv.ParseUint(pathParameters["userID"], 10, 64)
	if error != nil {
		responses.JsonResponse(w, http.StatusBadRequest, responses.ResponseError{Error: error.Error()})
		return
	}

	url := fmt.Sprintf("%s/users/%d/follow", config.ApiUrl, userId)

	response, error := requests.RequestAuthenticatedEndpoint(r, http.MethodPost, url, nil)
	if error != nil {
		responses.JsonResponse(w, http.StatusInternalServerError, responses.ResponseError{Error: error.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.ErrorResponseResolver(w, response)
		return
	}
	responses.JsonResponse(w, 204, nil) // fixed 204 for jquery not require a json body as response
}

// UnfollowUser deixar de seguir um usuário
func UnfollowUser(w http.ResponseWriter, r *http.Request) {

	pathParameters := mux.Vars(r)
	userId, error := strconv.ParseUint(pathParameters["userID"], 10, 64)
	if error != nil {
		responses.JsonResponse(w, http.StatusBadRequest, responses.ResponseError{Error: error.Error()})
		return
	}

	url := fmt.Sprintf("%s/users/%d/unfollow", config.ApiUrl, userId)

	response, error := requests.RequestAuthenticatedEndpoint(r, http.MethodPost, url, nil)
	if error != nil {
		responses.JsonResponse(w, http.StatusInternalServerError, responses.ResponseError{Error: error.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.ErrorResponseResolver(w, response)
		return
	}
	responses.JsonResponse(w, 204, nil) // fixed 204 for jquery not require a json body as response
}

// UpdateUser calls API to update a user
func UpdateUser(w http.ResponseWriter, r *http.Request) {

	cookie, _ := security.ReadAuthCookie(r)
	loggedUserId, _ := strconv.ParseUint(cookie["id"], 10, 64)

	r.ParseForm()

	user, error := json.Marshal(map[string]string{
		"name": r.FormValue("name"),
		"nick": r.FormValue("nick"),
		"email": r.FormValue("email"),
	})
	if error != nil {
		responses.JsonResponse(w, http.StatusBadRequest, responses.ResponseError{Error: error.Error()})
		return
	}

	url := fmt.Sprintf("%s/users/%d", config.ApiUrl, loggedUserId)
	response, error := requests.RequestAuthenticatedEndpoint(r, http.MethodPut, url, bytes.NewBuffer(user))
	if error != nil {
		responses.JsonResponse(w, http.StatusInternalServerError, responses.ResponseError{Error: error.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.ErrorResponseResolver(w, response)
		return
	}
	responses.JsonResponse(w, 204, nil) // fixed 204 for jquery not require a json body as response
}

// ChangePassword calls API to change a user's password
func ChangePassword(w http.ResponseWriter, r *http.Request) {

	cookie, _ := security.ReadAuthCookie(r)
	loggedUserId, _ := strconv.ParseUint(cookie["id"], 10, 64)

	r.ParseForm()

	password, error := json.Marshal(map[string]string{
		"previous": r.FormValue("oldPassword"),
		"new": r.FormValue("newPassword"),
	})
	if error != nil {
		responses.JsonResponse(w, http.StatusBadRequest, responses.ResponseError{Error: error.Error()})
		return
	}

	url := fmt.Sprintf("%s/users/%d/update-password", config.ApiUrl, loggedUserId)
	response, error := requests.RequestAuthenticatedEndpoint(r, http.MethodPost, url, bytes.NewBuffer(password))
	if error != nil {
		responses.JsonResponse(w, http.StatusInternalServerError, responses.ResponseError{Error: error.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.ErrorResponseResolver(w, response)
		return
	}
	responses.JsonResponse(w, 204, nil) // fixed 204 for jquery not require a json body as response
}

// DeleteAccount delete de logged user account
func DeleteAccount(w http.ResponseWriter, r *http.Request) {

	cookie, _ := security.ReadAuthCookie(r)
	loggedUserId, _ := strconv.ParseUint(cookie["id"], 10, 64)

	url := fmt.Sprintf("%s/users/%d", config.ApiUrl, loggedUserId)
	response, error := requests.RequestAuthenticatedEndpoint(r, http.MethodDelete, url, nil)
	if error != nil {
		responses.JsonResponse(w, http.StatusInternalServerError, responses.ResponseError{Error: error.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.ErrorResponseResolver(w, response)
		return
	}

	security.DeleteCookieValue(w)
}