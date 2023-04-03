package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"web/src/responses"
)

// CreateUser do a request on API to create an user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	name := r.FormValue("name")
	email := r.FormValue("email")
	nick := r.FormValue("nick")
	password := r.FormValue("password")

	user, err := json.Marshal(map[string]string{
		"name":     name,
		"email":    email,
		"nick":     nick,
		"password": password,
	})
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{Err: err.Error()})
		return
	}

	response, err := http.Post(
		"http://localhost:5000/users",
		"application/json",
		bytes.NewBuffer(user),
	)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Err: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.FixStatusCodeError(w, response)
		return
	}

	responses.JSON(w, response.StatusCode, nil)
}
