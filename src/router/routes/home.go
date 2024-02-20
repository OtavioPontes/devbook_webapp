package routes

import (
	"devbook_webapp/src/controllers"
	"net/http"
)

var routeHome = Route{
	Uri:                   "/home",
	Method:                http.MethodGet,
	Function:              controllers.LoadHomePage,
	RequireAuthentication: true,
}
