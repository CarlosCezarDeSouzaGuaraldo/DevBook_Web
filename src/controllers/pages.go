package controllers

import (
	"fmt"
	"net/http"
	"web/src/config"
	"web/src/requests"
	"web/src/utils"
)

// LoadLoginScreen render the login screen
func LoadLoginScreen(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "login.html", nil)
}

// LoadSignupScreen render the signup screen
func LoadSignupScreen(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "signup.html", nil)
}

// LoadHomePage render the home screen
func LoadHomePage(w http.ResponseWriter, r *http.Request) {

	url := fmt.Sprintf("%s/publications", config.URL_API)
	response, err := requests.DoAuthRequest(r, http.MethodGet, url, nil)
	fmt.Println(response.StatusCode, err)

	utils.ExecuteTemplate(w, "home.html", nil)
}
