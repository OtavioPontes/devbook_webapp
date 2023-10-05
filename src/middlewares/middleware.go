package middlewares

import (
	"devbook_webapp/src/cookies"
	"log"
	"net/http"
)

func Logger(nextFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		nextFunc(w, r)
	}
}

func Authenticatefunc(nextFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := cookies.Read(r)
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusMovedPermanently)
			return
		}

		nextFunc(w, r)
	}
}
