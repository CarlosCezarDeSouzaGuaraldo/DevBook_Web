package controllers

import (
	"net/http"
	"web/src/cookies"
)

// DoLogout remove the user auth data
func DoLogout(w http.ResponseWriter, r *http.Request) {
	cookies.Remove(w)
	http.Redirect(w, r, "/login", 302)
}
