package routes

import (
	"devbook_webapp/src/controllers"
	"net/http"
)

var routesPosts = []Route{
	{
		Uri:                   "/posts",
		Method:                http.MethodPost,
		Function:              controllers.CreatePost,
		RequireAuthentication: true,
	},
	{
		Uri:                   "/posts/{postId}/like",
		Method:                http.MethodPost,
		Function:              controllers.LikePost,
		RequireAuthentication: true,
	},
	{
		Uri:                   "/posts/{postId}/dislike",
		Method:                http.MethodPost,
		Function:              controllers.DislikePost,
		RequireAuthentication: true,
	},
	{
		Uri:                   "/posts/{postId}/edit",
		Method:                http.MethodGet,
		Function:              controllers.LoadEditPostPage,
		RequireAuthentication: true,
	},
	{
		Uri:                   "/posts/{postId}",
		Method:                http.MethodPut,
		Function:              controllers.UpdatePost,
		RequireAuthentication: true,
	},
	{
		Uri:                   "/posts/{postId}",
		Method:                http.MethodDelete,
		Function:              controllers.DeletePost,
		RequireAuthentication: true,
	},
}
