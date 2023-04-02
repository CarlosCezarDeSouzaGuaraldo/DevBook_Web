package controllers

import (
	"net/http"
	"web/src/utils"
)

// LoadLoginScreen render the login screen
func LoadLoginScreen(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "login.html", nil)
}
