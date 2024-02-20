package controllers

import (
	"devbook_webapp/src/cookies"
	"net/http"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	cookies.Delete(w)
	http.Redirect(w, r, "/login", http.StatusFound)
}
