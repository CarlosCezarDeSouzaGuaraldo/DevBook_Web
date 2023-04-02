package controllers

import (
	"bytes"
	"encoding/json"
	"log"
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
		log.Fatal(err)
		return
	}

	response, err := http.Post(
		"http://localhost:5000/users",
		"application/json",
		bytes.NewBuffer(user),
	)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer response.Body.Close()

	responses.JSON(w, response.StatusCode, nil)
}
