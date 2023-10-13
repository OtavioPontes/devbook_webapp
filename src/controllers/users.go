package controllers

import (
	"bytes"
	"devbook_webapp/src/config"
	"devbook_webapp/src/cookies"
	"devbook_webapp/src/requests"
	"devbook_webapp/src/responses"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	user, err := json.Marshal(map[string]string{
		"name":     r.FormValue("name"),
		"email":    r.FormValue("email"),
		"nick":     r.FormValue("nick"),
		"password": r.FormValue("password"),
	})

	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{Error: err.Error()})
		return
	}

	response, err := http.Post(fmt.Sprintf("%s/users", config.ApiUrl), "application/json", bytes.NewBuffer(user))

	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.HandleStatusCodeErrors(w, response)
		return
	}

	responses.JSON(w, response.StatusCode, nil)
}

func Unfollow(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userId, err := strconv.ParseUint(params["id"], 10, 64)

	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{Error: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/users/%d/unfollow", config.ApiUrl, userId)

	response, err := requests.RequestWithAuth(r, http.MethodPost, url, nil)

	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: err.Error()})
		return
	}

	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.HandleStatusCodeErrors(w, response)
		return
	}

	responses.JSON(w, response.StatusCode, nil)

}

func Follow(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userId, err := strconv.ParseUint(params["id"], 10, 64)

	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{Error: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/users/%d/follow", config.ApiUrl, userId)

	response, err := requests.RequestWithAuth(r, http.MethodPost, url, nil)

	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: err.Error()})
		return
	}

	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.HandleStatusCodeErrors(w, response)
		return
	}

	responses.JSON(w, response.StatusCode, nil)

}

func EditProfile(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	user, err := json.Marshal(map[string]string{
		"name":  r.FormValue("name"),
		"nick":  r.FormValue("nick"),
		"email": r.FormValue("email"),
	})

	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{Error: err.Error()})
		return
	}

	cookie, _ := cookies.Read(r)
	userId, _ := strconv.ParseUint(cookie["id"], 10, 64)

	url := fmt.Sprintf("%s/users/%d", config.ApiUrl, userId)

	response, err := requests.RequestWithAuth(r, http.MethodPut, url, bytes.NewBuffer(user))
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: err.Error()})
		return
	}

	defer response.Body.Close()
	if response.StatusCode >= 400 {
		responses.HandleStatusCodeErrors(w, response)
		return
	}

	responses.JSON(w, response.StatusCode, nil)

}

func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	user, err := json.Marshal(map[string]string{
		"current": r.FormValue("password"),
		"new":     r.FormValue("newPassword"),
	})

	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{Error: err.Error()})
		return
	}

	cookie, _ := cookies.Read(r)
	userId, _ := strconv.ParseUint(cookie["id"], 10, 64)

	url := fmt.Sprintf("%s/users/%d/update-password", config.ApiUrl, userId)

	response, err := requests.RequestWithAuth(r, http.MethodPost, url, bytes.NewBuffer(user))
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: err.Error()})
		return
	}

	defer response.Body.Close()
	if response.StatusCode >= 400 {
		responses.HandleStatusCodeErrors(w, response)
		return
	}

	responses.JSON(w, response.StatusCode, nil)

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	cookie, _ := cookies.Read(r)
	userId, _ := strconv.ParseUint(cookie["id"], 10, 64)

	url := fmt.Sprintf("%s/users/%d", config.ApiUrl, userId)

	response, err := requests.RequestWithAuth(r, http.MethodDelete, url, nil)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: err.Error()})
		return
	}

	defer response.Body.Close()
	if response.StatusCode >= 400 {
		responses.HandleStatusCodeErrors(w, response)
		return
	}

	responses.JSON(w, response.StatusCode, nil)

}
