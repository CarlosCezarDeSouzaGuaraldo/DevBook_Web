package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"web/src/config"
	"web/src/requests"
	"web/src/responses"

	"github.com/gorilla/mux"
)

// CreatePublication send a request on API to create a publication
func CreatePublication(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	publication, err := json.Marshal(map[string]string{
		"title":   r.FormValue("title"),
		"content": r.FormValue("content"),
	})
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{Err: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/publications", config.URL_API)
	response, err := requests.DoAuthRequest(r, http.MethodPost, url, bytes.NewBuffer(publication))
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

// LikePublication send a request on API to like a publication
func LikePublication(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	publicationID, err := strconv.ParseUint(params["publicationId"], 10, 64)
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{Err: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/publications/%d/like", config.URL_API, publicationID)
	response, err := requests.DoAuthRequest(r, http.MethodPost, url, nil)
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

// UnlikePublication send a request on API to like a publication
func UnlikePublication(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	publicationID, err := strconv.ParseUint(params["publicationId"], 10, 64)
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{Err: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/publications/%d/unlike", config.URL_API, publicationID)
	response, err := requests.DoAuthRequest(r, http.MethodPost, url, nil)
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

// UpdatePublication send a request on API to update a publication
func UpdatePublication(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	publicationID, err := strconv.ParseUint(params["publicationId"], 10, 64)
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{Err: err.Error()})
		return
	}

	r.ParseForm()

	publication, err := json.Marshal(map[string]string{
		"title":   r.FormValue("title"),
		"content": r.FormValue("content"),
	})
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{Err: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/publications/%d", config.URL_API, publicationID)
	response, err := requests.DoAuthRequest(r, http.MethodPut, url, bytes.NewBuffer(publication))
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
