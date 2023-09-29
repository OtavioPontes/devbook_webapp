package routes

import (
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

		r.HandleFunc(route.Uri, route.Function).Methods(route.Method)

	}

	fileServer := http.FileServer(http.Dir("./assets/"))

	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return r

}
