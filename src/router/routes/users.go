package routes

import (
	"devbook_webapp/src/controllers"
	"net/http"
)

var routeUsers = []Route{
	{
		Uri:                   "/criar-usuario",
		Method:                http.MethodGet,
		Function:              controllers.LoadCreateUserPage,
		RequireAuthentication: false,
	},
	{
		Uri:                   "/users",
		Method:                http.MethodPost,
		Function:              controllers.CreateUser,
		RequireAuthentication: false,
	},
}
