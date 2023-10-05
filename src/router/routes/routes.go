package routes

import (
	"devbook_webapp/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Uri                   string
	Method                string
	Function              func(http.ResponseWriter, *http.Request)
	RequireAuthentication bool
}

func Configure(r *mux.Router) *mux.Router {

	var routes = routesLogin
	routes = append(routes, routeUsers...)
	routes = append(routes, routeHome)

	for _, route := range routes {
		if route.RequireAuthentication {
			r.HandleFunc(route.Uri, middlewares.Logger(middlewares.Authenticatefunc(route.Function))).Methods(route.Method)

		} else {
			r.HandleFunc(route.Uri, middlewares.Logger(route.Function)).Methods(route.Method)

		}
	}

	fileServer := http.FileServer(http.Dir("./assets/"))

	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return r

}
