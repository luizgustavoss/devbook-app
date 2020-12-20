package controllers

import (
	"devbookapp/src/utils"
	"net/http"
)

// LoadLoginPage loads login page
func LoadLoginPage(w http.ResponseWriter, r *http.Request) {

	utils.RenderTemplate(w, "login.html", nil)
}

// LoadCreateUserPage loads create user page
func LoadCreateUserPage(w http.ResponseWriter, r *http.Request) {

	utils.RenderTemplate(w, "create-user.html", nil)
}