package models

import (
	"devbook_webapp/src/config"
	"devbook_webapp/src/requests"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type User struct {
	Id         uint64    `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Nick       string    `json:"nick"`
	CreatedAt  time.Time `json:"created_at"`
	Followers  []User    `json:"followers"`
	Followings []User    `json:"followings"`
	Posts      []Post    `json:"posts"`
}

func GetUserFull(userId uint64, r *http.Request) (User, error) {
	channelUser := make(chan User)
	channelFollowers := make(chan []User)
	channelFollowings := make(chan []User)
	channelPosts := make(chan []Post)

	go GetUserData(channelUser, userId, r)
	go GetFollowersData(channelFollowers, userId, r)
	go GetFollowingsData(channelFollowings, userId, r)
	go GetPostsData(channelPosts, userId, r)

	var (
		user       User
		followers  []User
		followings []User
		posts      []Post
	)

	for i := 0; i < 4; i++ {
		select {
		case userLoaded := <-channelUser:
			if userLoaded.Id == 0 {
				return User{}, errors.New("falha ao carregar usuário")
			}
			user = userLoaded
		case followersLoaded := <-channelFollowers:
			if followersLoaded == nil {
				return User{}, errors.New("falha ao carregar seguidores")
			}
			followers = followersLoaded
		case followingsLoaded := <-channelFollowings:
			if followingsLoaded == nil {
				return User{}, errors.New("falha ao carregar quem o usuário segue")
			}
			followings = followingsLoaded
		case postsLoaded := <-channelPosts:
			if postsLoaded == nil {
				return User{}, errors.New("falha ao carregar postagens")
			}
			posts = postsLoaded
		}
	}

	return User{
		Id:         user.Id,
		Name:       user.Name,
		Email:      user.Email,
		Nick:       user.Nick,
		CreatedAt:  user.CreatedAt,
		Followers:  followers,
		Followings: followings,
		Posts:      posts,
	}, nil
}

func GetUserData(channel chan<- User, userId uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d", config.ApiUrl, userId)

	response, err := requests.RequestWithAuth(r, http.MethodGet, url, nil)

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

func GetFollowersData(channel chan<- []User, userId uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d/followers", config.ApiUrl, userId)

	response, err := requests.RequestWithAuth(r, http.MethodGet, url, nil)

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

func GetFollowingsData(channel chan<- []User, userId uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d/followings", config.ApiUrl, userId)

	response, err := requests.RequestWithAuth(r, http.MethodGet, url, nil)

	if err != nil {
		channel <- nil
		return
	}

	defer response.Body.Close()

	var followings []User

	if err = json.NewDecoder(response.Body).Decode(&followings); err != nil {
		channel <- nil
		return
	}

	if followings == nil {
		channel <- make([]User, 0)
		return
	}

	channel <- followings
}

func GetPostsData(channel chan<- []Post, userId uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d/posts", config.ApiUrl, userId)

	response, err := requests.RequestWithAuth(r, http.MethodGet, url, nil)

	if err != nil {
		channel <- nil
		return
	}

	defer response.Body.Close()

	var posts []Post

	if err = json.NewDecoder(response.Body).Decode(&posts); err != nil {
		channel <- nil
		return
	}

	if posts == nil {
		channel <- make([]Post, 0)
		return
	}

	channel <- posts
}
