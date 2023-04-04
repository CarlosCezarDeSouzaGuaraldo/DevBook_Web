package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"web/src/config"
	"web/src/cookies"
	"web/src/models"
	"web/src/responses"
)

// DoLogin use the e-mail and password to login the user
func DoLogin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	email := r.FormValue("email")
	password := r.FormValue("password")

	user, err := json.Marshal(map[string]string{
		"email":    email,
		"password": password,
	})
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{Err: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/login", config.URL_API)

	response, err := http.Post(
		url,
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

	var authData models.AuthData
	if err = json.NewDecoder(response.Body).Decode(&authData); err != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.ErrorAPI{Err: err.Error()})
		return
	}

	if err = cookies.Save(w, authData.ID, authData.Token); err != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.ErrorAPI{Err: err.Error()})
		return
	}

	responses.JSON(w, http.StatusOK, nil)
}
