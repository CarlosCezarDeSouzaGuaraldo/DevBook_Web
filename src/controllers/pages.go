package controllers

import (
	"net/http"
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
	
	utils.ExecuteTemplate(w, "home.html", nil)
}
