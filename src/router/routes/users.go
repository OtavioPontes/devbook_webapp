package routes

import (
	"devbook_webapp/src/controllers"
	"net/http"
)

var routeUsers = []Route{
	{
		Uri:                   "/create-user",
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
	{
		Uri:                   "/get-users",
		Method:                http.MethodGet,
		Function:              controllers.LoadUsersPage,
		RequireAuthentication: true,
	},
	{
		Uri:                   "/users/{id}",
		Method:                http.MethodGet,
		Function:              controllers.LoadUserPage,
		RequireAuthentication: true,
	},
	{
		Uri:                   "/users/{id}/unfollow",
		Method:                http.MethodPost,
		Function:              controllers.Unfollow,
		RequireAuthentication: true,
	},
	{
		Uri:                   "/users/{id}/follow",
		Method:                http.MethodPost,
		Function:              controllers.Follow,
		RequireAuthentication: true,
	},
	{
		Uri:                   "/profile",
		Method:                http.MethodGet,
		Function:              controllers.LoadProfilePage,
		RequireAuthentication: true,
	},
	{
		Uri:                   "/edit-user",
		Method:                http.MethodGet,
		Function:              controllers.LoadEditProfilePage,
		RequireAuthentication: true,
	},
	{
		Uri:                   "/edit-profile",
		Method:                http.MethodPut,
		Function:              controllers.EditProfile,
		RequireAuthentication: true,
	},
}
