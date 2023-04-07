package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
	"web/src/config"
	"web/src/requests"
)

// User is a user from social media
type User struct {
	ID           uint64         `json:"id,omitempty"`
	Name         string         `json:"name,omitempty"`
	Nick         string         `json:"nick,omitempty"`
	Email        string         `json:"email,omitempty"`
	CreatedAt    time.Time      `json:"createdAt,omitempty"`
	Followers    []User         `json:"followers,omitempty"`
	Following    []User         `json:"following,omitempty"`
	Publications []Publications `json:"publications,omitempty"`
}

// FindCompleteUser call the API to get followers, following and publications
func FindCompleteUser(userID uint64, r *http.Request) (User, error) {
	userChannel := make(chan User)
	followersChannel := make(chan []User)
	followingChannel := make(chan []User)
	publicationsChannel := make(chan []Publications)

	go GetUserData(userChannel, userID, r)
	go GetFollowersData(followersChannel, userID, r)
	go GetFollowingData(followingChannel, userID, r)
	go GetPublicationsData(publicationsChannel, userID, r)

	var (
		user         User
		followers    []User
		following    []User
		publications []Publications
	)

	for i := 0; i < 4; i++ {
		select {
		case loadedUser := <-userChannel:
			if loadedUser.ID == 0 {
				return User{}, errors.New("Error to find user")
			}
			user = loadedUser
		case loadedFollowers := <-followersChannel:
			if loadedFollowers == nil {
				return User{}, errors.New("Error to find followers")
			}
			followers = loadedFollowers
		case loadedFollowing := <-followingChannel:
			if loadedFollowing == nil {
				return User{}, errors.New("Error to find following users")
			}
			following = loadedFollowing
		case loadedPublications := <-publicationsChannel:
			if loadedPublications == nil {
				return User{}, errors.New("Error to find user's publications")
			}
			publications = loadedPublications
		}
	}

	user.Followers = followers
	user.Following = following
	user.Publications = publications

	return user, nil
}

// GetUserData do a request on API to get user
func GetUserData(channel chan<- User, userID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d", config.URL_API, userID)
	response, err := requests.DoAuthRequest(r, http.MethodGet, url, nil)
	if err != nil {
		channel <- User{}
		return
	}
	defer response.Body.Close()

	var user User
	if err = json.NewDecoder(response.Body).Decode(&user); err != nil {
		channel <- User{}
		return
	}

	channel <- user
}

// GetFollowersData do a request on API to get the user's followers
func GetFollowersData(channel chan<- []User, userID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d/followers", config.URL_API, userID)
	response, err := requests.DoAuthRequest(r, http.MethodGet, url, nil)
	if err != nil {
		channel <- nil
		return
	}
	defer response.Body.Close()

	var followers []User
	if err = json.NewDecoder(response.Body).Decode(&followers); err != nil {
		channel <- nil
		return
	}

	if followers == nil {
		channel <- make([]User, 0)
		return
	}

	channel <- followers
}

// GetFollowingData do a request on API to get the user's following users
func GetFollowingData(channel chan<- []User, userID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d/following", config.URL_API, userID)
	response, err := requests.DoAuthRequest(r, http.MethodGet, url, nil)
	if err != nil {
		channel <- nil
		return
	}
	defer response.Body.Close()

	var following []User
	if err = json.NewDecoder(response.Body).Decode(&following); err != nil {
		channel <- nil
		return
	}

	if following == nil {
		channel <- make([]User, 0)
		return
	}

	channel <- following
}

// GetPublicationsData do a request on API to get the user's publications
func GetPublicationsData(channel chan<- []Publications, userID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d/publications", config.URL_API, userID)
	response, err := requests.DoAuthRequest(r, http.MethodGet, url, nil)
	if err != nil {
		channel <- nil
		return
	}
	defer response.Body.Close()

	var publications []Publications
	if err = json.NewDecoder(response.Body).Decode(&publications); err != nil {
		channel <- nil
		return
	}

	if publications == nil {
		channel <- make([]Publications, 0)
		return
	}

	channel <- publications
}
