package controllers

import (
	"devbookapp/src/config"
	"devbookapp/src/models"
	"devbookapp/src/requests"
	"devbookapp/src/responses"
	"devbookapp/src/security"
	"devbookapp/src/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"strings"
)

// LoadLoginPage loads login page
func LoadLoginPage(w http.ResponseWriter, r *http.Request) {

	cookie, _ := security.ReadAuthCookie(r)
	if cookie["token"] != "" {
		http.Redirect(w, r, "/home", 302)
		return
	}
	utils.RenderTemplate(w, "login.html", nil)
}

// LoadCreateUserPage loads create user page
func LoadCreateUserPage(w http.ResponseWriter, r *http.Request) {

	utils.RenderTemplate(w, "create-user.html", nil)
}

// LoadHomePage loads home page
func LoadHomePage(w http.ResponseWriter, r *http.Request) {

	url := fmt.Sprintf("%s/publications", config.ApiUrl)
	response, error := requests.RequestAuthenticatedEndpoint(r, http.MethodGet, url, nil)

	if error != nil {
		responses.JsonResponse(w, http.StatusInternalServerError, responses.ResponseError{Error: error.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.ErrorResponseResolver(w, response)
		return
	}

	var publications []models.Publication
	if error = json.NewDecoder(response.Body).Decode(&publications); error != nil {
		responses.JsonResponse(w, http.StatusUnprocessableEntity, responses.ResponseError{Error: error.Error()})
		return
	}

	cookies, _ := security.ReadAuthCookie(r)

	loggedUserId, _  := strconv.ParseUint(cookies["id"], 10, 64)

	utils.RenderTemplate(w, "home.html", struct{
		Publications []models.Publication
		LoggedUserID uint64
	}{
		Publications: publications,
		LoggedUserID: loggedUserId,
	})
}


// LoadPublicationDetail load details of a publication for edit
func LoadPublicationDetail(w http.ResponseWriter, r *http.Request) {

	pathParameters := mux.Vars(r)
	publicationId, error := strconv.ParseUint(pathParameters["publicationId"], 10, 64)
	if error != nil {
		responses.JsonResponse(w, http.StatusBadRequest, responses.ResponseError{Error: error.Error()})
		return
	}

	url := fmt.Sprintf("%s/publications/%d", config.ApiUrl, publicationId)

	response, error := requests.RequestAuthenticatedEndpoint(r, http.MethodGet, url, nil)

	if error != nil {
		responses.JsonResponse(w, http.StatusInternalServerError, responses.ResponseError{Error: error.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.ErrorResponseResolver(w, response)
		return
	}

	var publication models.Publication
	if error = json.NewDecoder(response.Body).Decode(&publication); error != nil {
		responses.JsonResponse(w, http.StatusUnprocessableEntity, responses.ResponseError{Error: error.Error()})
		return
	}

	utils.RenderTemplate(w, "edit-publication.html", publication)
}

// SearchUsers searches users given a name or nick
func SearchUsers(w http.ResponseWriter, r *http.Request) {

	nickOrName := strings.ToLower(r.URL.Query().Get("user"))

	url := fmt.Sprintf("%s/users?desc=%s", config.ApiUrl, nickOrName)

	response, error := requests.RequestAuthenticatedEndpoint(r, http.MethodGet, url, nil)

	if error != nil {
		responses.JsonResponse(w, http.StatusInternalServerError, responses.ResponseError{Error: error.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.ErrorResponseResolver(w, response)
		return
	}

	var users []models.User
	if error = json.NewDecoder(response.Body).Decode(&users); error != nil {
		responses.JsonResponse(w, http.StatusUnprocessableEntity, responses.ResponseError{Error: error.Error()})
		return
	}

	utils.RenderTemplate(w, "users.html", users)
}

// GetUserDetails loads user details and render page
func GetUserDetails(w http.ResponseWriter, r *http.Request) {

	pathParameters := mux.Vars(r)
	userId, error := strconv.ParseUint(pathParameters["userID"], 10, 64)
	if error != nil {
		responses.JsonResponse(w, http.StatusBadRequest, responses.ResponseError{Error: error.Error()})
		return
	}

	cookie, _ := security.ReadAuthCookie(r)
	loggedUserId, _ := strconv.ParseUint(cookie["id"], 10, 64)

	if loggedUserId == userId {
		http.Redirect(w, r, "profile", 302)
		return
	}

	user, error := models.LoadUserDetails(userId, r)
	if error != nil {
		responses.JsonResponse(w, http.StatusInternalServerError, responses.ResponseError{Error: error.Error()})
		return
	}

	utils.RenderTemplate(w, "user.html", struct{
		User models.User
		LoggedUserID uint64
	}{
		User: user,
		LoggedUserID: loggedUserId,
	})
}


// LoadLoggedUserProfile loads logged user details and render page
func LoadLoggedUserProfile(w http.ResponseWriter, r *http.Request) {

	cookie, _ := security.ReadAuthCookie(r)
	loggedUserId, _ := strconv.ParseUint(cookie["id"], 10, 64)

	user, error := models.LoadUserDetails(loggedUserId, r)
	if error != nil {
		responses.JsonResponse(w, http.StatusInternalServerError, responses.ResponseError{Error: error.Error()})
		return
	}
	utils.RenderTemplate(w, "profile.html", user)
}

// LoadEditUserPage renders page to change user details
func LoadEditUserPage(w http.ResponseWriter, r *http.Request) {

	cookie, _ := security.ReadAuthCookie(r)
	loggedUserId, _ := strconv.ParseUint(cookie["id"], 10, 64)

	url := fmt.Sprintf("%s/users/%d", config.ApiUrl, loggedUserId)

	response, error := requests.RequestAuthenticatedEndpoint(r, http.MethodGet, url, nil)

	if error != nil {
		responses.JsonResponse(w, http.StatusInternalServerError, responses.ResponseError{Error: error.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.ErrorResponseResolver(w, response)
		return
	}

	var user models.User
	if error = json.NewDecoder(response.Body).Decode(&user); error != nil {
		responses.JsonResponse(w, http.StatusUnprocessableEntity, responses.ResponseError{Error: error.Error()})
		return
	}

	utils.RenderTemplate(w, "edit-user.html", user)

}

// LoadChangePasswordPage renders page to change user password
func LoadChangePasswordPage(w http.ResponseWriter, r *http.Request) {

	cookie, _ := security.ReadAuthCookie(r)
	loggedUserId, _ := strconv.ParseUint(cookie["id"], 10, 64)

	url := fmt.Sprintf("%s/users/%d", config.ApiUrl, loggedUserId)

	response, error := requests.RequestAuthenticatedEndpoint(r, http.MethodGet, url, nil)

	if error != nil {
		responses.JsonResponse(w, http.StatusInternalServerError, responses.ResponseError{Error: error.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.ErrorResponseResolver(w, response)
		return
	}

	var user models.User
	if error = json.NewDecoder(response.Body).Decode(&user); error != nil {
		responses.JsonResponse(w, http.StatusUnprocessableEntity, responses.ResponseError{Error: error.Error()})
		return
	}

	utils.RenderTemplate(w, "change-password.html", user)
}