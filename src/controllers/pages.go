package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
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
	cookie, _ := cookies.Read(r)
	if cookie["token"] != "" {
		http.Redirect(w, r, "/home", 302)
		return
	}

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

// LoadUsersPage render the users list screen
func LoadUsersPage(w http.ResponseWriter, r *http.Request) {
	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))

	url := fmt.Sprintf("%s/users?user=%s", config.URL_API, nameOrNick)

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

	var users []models.User
	if err = json.NewDecoder(response.Body).Decode(&users); err != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.ErrorAPI{Err: err.Error()})
		return
	}

	utils.ExecuteTemplate(w, "users.html", users)
}

// LoadProfileUser do a request on API to get an user
func LoadProfileUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{Err: err.Error()})
		return
	}

	cookie, _ := cookies.Read(r)
	userLogged, _ := strconv.ParseUint(cookie["id"], 10, 64)

	if userID == userLogged {
		http.Redirect(w, r, "/profile", 302)
	}

	user, err := models.FindCompleteUser(userID, r)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Err: err.Error()})
		return
	}

	utils.ExecuteTemplate(w, "user.html", struct {
		User         models.User
		UserLoggedID uint64
	}{
		User:         user,
		UserLoggedID: userLogged,
	})
}

// LoadLoggedProfileUser do a request on API to get the logged user
func LoadLoggedProfileUser(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Read(r)
	userLogged, _ := strconv.ParseUint(cookie["id"], 10, 64)

	user, err := models.FindCompleteUser(userLogged, r)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Err: err.Error()})
		return
	}

	utils.ExecuteTemplate(w, "profile.html", user)
}

// LoadEditProfileUserScreen load the profile user screen
func LoadEditProfileUserScreen(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Read(r)
	userLoggedID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	channel := make(chan models.User)
	go models.GetUserData(channel, userLoggedID, r)
	user := <-channel

	if user.ID == 0 {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Err: "Error to find user"})
		return
	}

	utils.ExecuteTemplate(w, "edit-user.html", user)
}

// LoadUpdatePasswordScreen load the update user password screen
func LoadUpdatePasswordScreen(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "update-password.html", nil)
}
