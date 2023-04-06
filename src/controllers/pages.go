package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"web/src/config"
	"web/src/cookies"
	"web/src/models"
	"web/src/requests"
	"web/src/responses"
	"web/src/utils"

	"github.com/gorilla/mux"
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
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Err: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.FixStatusCodeError(w, response)
		return
	}

	var publications []models.Publications
	if err = json.NewDecoder(response.Body).Decode(&publications); err != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.ErrorAPI{Err: err.Error()})
		return
	}

	cookie, _ := cookies.Read(r)
	userID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	utils.ExecuteTemplate(w, "home.html", struct {
		Publications []models.Publications
		UserID       uint64
	}{
		Publications: publications,
		UserID:       userID,
	})
}

// LoadEditPublicationScreen render the edit publication screen
func LoadEditPublicationScreen(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	publicationID, err := strconv.ParseUint(params["publicationId"], 10, 64)
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{Err: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/publications/%d", config.URL_API, publicationID)
	response, err := requests.DoAuthRequest(r, http.MethodGet, url, nil)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Err: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.FixStatusCodeError(w, response)
		return
	}

	var publication models.Publications
	if err = json.NewDecoder(response.Body).Decode(&publication); err != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.ErrorAPI{Err: err.Error()})
		return
	}

	utils.ExecuteTemplate(w, "update-publication.html", publication)
}
