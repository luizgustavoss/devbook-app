package controllers

import (
	"bytes"
	"devbookapp/src/config"
	"devbookapp/src/requests"
	"devbookapp/src/responses"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// CreatePublication calls API to create a publication
func CreatePublication(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	publication, error := json.Marshal(map[string]string{
		"title": r.FormValue("title"),
		"content": r.FormValue("content"),
	})
	if error != nil {
		responses.JsonResponse(w, http.StatusBadRequest, responses.ResponseError{Error: error.Error()})
		return
	}

	url := fmt.Sprintf("%s/publications", config.ApiUrl)

	response, error := requests.RequestAuthenticatedEndpoint(r, http.MethodPost, url, bytes.NewBuffer(publication))

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


// LikePublication register like in a publication
func LikePublication(w http.ResponseWriter, r *http.Request) {

	pathParameters := mux.Vars(r)
	publicationId, error := strconv.ParseUint(pathParameters["publicationId"], 10, 64)
	if error != nil {
		responses.JsonResponse(w, http.StatusBadRequest, responses.ResponseError{Error: error.Error()})
		return
	}

	url := fmt.Sprintf("%s/publications/%d/like", config.ApiUrl, publicationId)

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


// UnlikePublication register unlike in a publication
func UnlikePublication(w http.ResponseWriter, r *http.Request) {

	pathParameters := mux.Vars(r)
	publicationId, error := strconv.ParseUint(pathParameters["publicationId"], 10, 64)
	if error != nil {
		responses.JsonResponse(w, http.StatusBadRequest, responses.ResponseError{Error: error.Error()})
		return
	}

	url := fmt.Sprintf("%s/publications/%d/unlike", config.ApiUrl, publicationId)

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


// DeletePublication deletes a publication
func DeletePublication(w http.ResponseWriter, r *http.Request) {

	pathParameters := mux.Vars(r)
	publicationId, error := strconv.ParseUint(pathParameters["publicationId"], 10, 64)
	if error != nil {
		responses.JsonResponse(w, http.StatusBadRequest, responses.ResponseError{Error: error.Error()})
		return
	}

	url := fmt.Sprintf("%s/publications/%d", config.ApiUrl, publicationId)

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

	responses.JsonResponse(w, 204, nil) // fixed 204 for jquery not require a json body as response

}

// UpdatePublication updates a publication
func UpdatePublication(w http.ResponseWriter, r *http.Request) {

	pathParameters := mux.Vars(r)
	publicationId, error := strconv.ParseUint(pathParameters["publicationId"], 10, 64)
	if error != nil {
		responses.JsonResponse(w, http.StatusBadRequest, responses.ResponseError{Error: error.Error()})
		return
	}

	r.ParseForm()

	publication, error := json.Marshal(map[string]string{
		"title": r.FormValue("title"),
		"content": r.FormValue("content"),
	})
	if error != nil {
		responses.JsonResponse(w, http.StatusBadRequest, responses.ResponseError{Error: error.Error()})
		return
	}

	url := fmt.Sprintf("%s/publications/%d", config.ApiUrl, publicationId)

	response, error := requests.RequestAuthenticatedEndpoint(r, http.MethodPut, url, bytes.NewBuffer(publication))

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