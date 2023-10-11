package routes

import (
	"devbook_webapp/src/controllers"
	"net/http"
)

var routeLogout = Route{
	Uri:                   "/logout",
	Method:                http.MethodGet,
	Function:              controllers.Logout,
	RequireAuthentication: true,
}
