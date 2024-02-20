package routes

import (
	"devbook_webapp/src/controllers"
	"net/http"
)

var routesLogin = []Route{
	{
		Uri:                   "/",
		Method:                http.MethodGet,
		Function:              controllers.LoadLogin,
		RequireAuthentication: false,
	},
	{
		Uri:                   "/login",
		Method:                http.MethodGet,
		Function:              controllers.LoadLogin,
		RequireAuthentication: false,
	},
	{
		Uri:                   "/login",
		Method:                http.MethodPost,
		Function:              controllers.Login,
		RequireAuthentication: false,
	},
}
