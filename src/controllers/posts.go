package controllers

import (
	"bytes"
	"devbook_webapp/src/config"
	"devbook_webapp/src/requests"
	"devbook_webapp/src/responses"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	post, err := json.Marshal(map[string]string{
		"title":   r.FormValue("title"),
		"content": r.FormValue("content"),
	})

	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{Error: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/posts", config.ApiUrl)

	response, err := requests.RequestWithAuth(r, http.MethodPost, url, bytes.NewBuffer(post))

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

func LikePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	postId, err := strconv.ParseUint(params["postId"], 10, 64)
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{Error: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/posts/%d/like", config.ApiUrl, postId)

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

func DislikePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	postId, err := strconv.ParseUint(params["postId"], 10, 64)
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{Error: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/posts/%d/dislike", config.ApiUrl, postId)

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

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	postId, err := strconv.ParseUint(params["postId"], 10, 64)

	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{Error: err.Error()})
		return
	}

	r.ParseForm()

	post, err := json.Marshal(map[string]string{
		"title":   r.FormValue("title"),
		"content": r.FormValue("content"),
	})

	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{Error: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/posts/%d", config.ApiUrl, postId)

	response, err := requests.RequestWithAuth(r, http.MethodPut, url, bytes.NewBuffer(post))

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

func DeletePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	postId, err := strconv.ParseUint(params["postId"], 10, 64)

	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{Error: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/posts/%d", config.ApiUrl, postId)

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
