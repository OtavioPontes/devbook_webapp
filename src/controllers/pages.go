package controllers

import (
	"devbook_webapp/src/config"
	"devbook_webapp/src/cookies"
	"devbook_webapp/src/models"
	"devbook_webapp/src/requests"
	"devbook_webapp/src/responses"
	"devbook_webapp/src/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func LoadLogin(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Read(r)
	if cookie["token"] != "" {
		http.Redirect(w, r, "/home", http.StatusFound)
	}
	utils.ExecuteTemplate(w, "login.html", nil)
}

func LoadCreateUserPage(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "register.html", nil)
}

func LoadHomePage(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/posts", config.ApiUrl)

	response, err := requests.RequestWithAuth(r, http.MethodGet, url, nil)

	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.HandleStatusCodeErrors(w, response)
		return
	}

	var posts []models.Post

	if err = json.NewDecoder(response.Body).Decode(&posts); err != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.ErrorAPI{Error: err.Error()})
		return
	}

	cookie, _ := cookies.Read(r)

	userId, _ := strconv.ParseUint(cookie["id"], 10, 64)

	utils.ExecuteTemplate(w, "home.html", struct {
		Posts  []models.Post
		UserId uint64
	}{
		Posts:  posts,
		UserId: userId,
	})
}

func LoadEditPostPage(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	postId, err := strconv.ParseUint(params["postId"], 10, 64)

	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{Error: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/posts/%d", config.ApiUrl, postId)

	response, err := requests.RequestWithAuth(r, http.MethodGet, url, nil)

	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.HandleStatusCodeErrors(w, response)
		return
	}

	var post models.Post

	if err = json.NewDecoder(response.Body).Decode(&post); err != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.ErrorAPI{Error: err.Error()})
		return
	}

	utils.ExecuteTemplate(w, "edit-post.html", post)
}

func LoadUsersPage(w http.ResponseWriter, r *http.Request) {
	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))
	url := fmt.Sprintf("%s/users?user=%s", config.ApiUrl, nameOrNick)

	response, err := requests.RequestWithAuth(r, http.MethodGet, url, nil)

	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.HandleStatusCodeErrors(w, response)
		return
	}

	var users []models.User

	if err := json.NewDecoder(response.Body).Decode(&users); err != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.ErrorAPI{Error: err.Error()})
		return
	}

	utils.ExecuteTemplate(w, "users.html", users)
}

func LoadUserPage(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userId, err := strconv.ParseUint(params["id"], 10, 64)

	cookie, _ := cookies.Read(r)
	userLoggedId, _ := strconv.ParseUint(cookie["id"], 10, 64)

	if userLoggedId == userId {
		http.Redirect(w, r, "/profile", http.StatusFound)
		return
	}

	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{Error: err.Error()})
		return
	}

	user, err := models.GetUserFull(userId, r)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: err.Error()})
		return
	}

	utils.ExecuteTemplate(w, "user.html", struct {
		User   models.User
		UserId uint64
	}{
		User:   user,
		UserId: userLoggedId,
	})

}

func LoadProfilePage(w http.ResponseWriter, r *http.Request) {

	cookie, _ := cookies.Read(r)
	userId, _ := strconv.ParseUint(cookie["id"], 10, 64)

	user, err := models.GetUserFull(userId, r)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: err.Error()})
		return
	}

	utils.ExecuteTemplate(w, "profile.html", user)

}

func LoadEditProfilePage(w http.ResponseWriter, r *http.Request) {

	cookie, _ := cookies.Read(r)
	userId, _ := strconv.ParseUint(cookie["id"], 10, 64)

	channelUser := make(chan models.User)

	go models.GetUserData(channelUser, userId, r)

	user := <-channelUser

	if user.Id == 0 {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: "Erro ao buscar o usuário"})
		return
	}

	utils.ExecuteTemplate(w, "edit-profile.html", user)

}

func LoadUpdatePasswordPage(w http.ResponseWriter, r *http.Request) {

	utils.ExecuteTemplate(w, "edit-password.html", nil)

}
