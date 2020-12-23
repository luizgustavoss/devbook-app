package models

import (
	"devbookapp/src/config"
	"devbookapp/src/requests"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type User struct {
	ID           uint64        `json:"id,omitempty"`
	Name         string        `json:"name,omitempty"`
	Nick         string        `json:"nick,omitempty"`
	Email        string        `json:"email,omitempty"`
	CreatedAt    time.Time     `json:"createdAt,omitempty"`
	Followers    []User        `json:"followers,omitempty"`
	Following    []User        `json:"following,omitempty"`
	Publications []Publication `json:"publications,omitempty"`
}

// LoadUserDetails loads full user details making various API requests
func LoadUserDetails(userID uint64, r *http.Request) (User, error) {

	userChannel := make(chan User)
	followersChannel := make(chan []User)
	followingChannel := make(chan []User)
	publicationsChannel := make(chan []Publication)

	go loadUserDetails(userChannel, userID, r)
	go loadUserFollowers(followersChannel, userID, r)
	go loadUserFollowing(followingChannel, userID, r)
	go loadUserPublications(publicationsChannel, userID, r)

	var (
		user User
		followers []User
		following []User
		publications []Publication
	)

	for i := 0; i < 4; i++ {
		select {
		case returnedUser := <-userChannel:{}
			if returnedUser.ID == 0 {
				return User{}, errors.New("Falha ao recuperar detalhes do usuário!")
			}
			user = returnedUser
		case returnedFollowers := <-followersChannel:
			if returnedFollowers == nil {
				return User{}, errors.New("Falha ao recuperar seguidores do usuário!")
			}
			followers = returnedFollowers
		case returnedFollowing := <-followingChannel:
			if returnedFollowing == nil {
				return User{}, errors.New("Falha ao recuperar usuários seguidos pelo usuário!")
			}
			following = returnedFollowing
		case returnedPublications := <-publicationsChannel:
			if returnedPublications == nil {
				return User{}, errors.New("Falha ao recuperar publicações do usuário!")
			}
			publications = returnedPublications
		}
	}

	user.Publications = publications
	user.Followers = followers
	user.Following = following

	return user, nil

}

func loadUserDetails(channel chan<- User, userID uint64, r *http.Request) {

	url := fmt.Sprintf("%s/users/%d", config.ApiUrl, userID)

	response, error := requests.RequestAuthenticatedEndpoint(r, http.MethodGet, url, nil)
	if error != nil {
		channel <- User{}
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		channel <- User{}
	}

	var user User
	if error = json.NewDecoder(response.Body).Decode(&user); error != nil {
		channel <- User{}
	}

	channel <- user
}

func loadUserFollowers(channel chan<- []User, userID uint64, r *http.Request) {

	url := fmt.Sprintf("%s/users/%d/followers", config.ApiUrl, userID)

	response, error := requests.RequestAuthenticatedEndpoint(r, http.MethodGet, url, nil)
	if error != nil {
		channel <- nil
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		channel <- nil
	}

	var followers []User
	if error = json.NewDecoder(response.Body).Decode(&followers); error != nil {
		channel <- nil
	}

	if followers == nil {
		channel <- make([]User, 0)
		return
	}

	channel <- followers
}

func loadUserFollowing(channel chan<- []User, userID uint64, r *http.Request) {

	url := fmt.Sprintf("%s/users/%d/followed", config.ApiUrl, userID)

	response, error := requests.RequestAuthenticatedEndpoint(r, http.MethodGet, url, nil)
	if error != nil {
		channel <- nil
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		channel <- nil
	}

	var followed []User
	if error = json.NewDecoder(response.Body).Decode(&followed); error != nil {
		channel <- nil
	}

	if followed == nil {
		channel <- make([]User, 0)
		return
	}

	channel <- followed
}

func loadUserPublications(channel chan<- []Publication, userID uint64, r *http.Request) {

	url := fmt.Sprintf("%s/users/%d/publications", config.ApiUrl, userID)

	response, error := requests.RequestAuthenticatedEndpoint(r, http.MethodGet, url, nil)
	if error != nil {
		channel <- nil
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		channel <- nil
	}

	var publications []Publication
	if error = json.NewDecoder(response.Body).Decode(&publications); error != nil {
		channel <- nil
	}

	if publications == nil {
		channel <- make([]Publication, 0)
		return
	}

	channel <- publications
}
