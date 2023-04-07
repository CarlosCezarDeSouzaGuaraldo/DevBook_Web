package routes

import (
	"net/http"
	"web/src/controllers"
)

var userRoutes = []Route{
	{
		URI:            "/signup",
		Method:         http.MethodGet,
		Func:           controllers.LoadSignupScreen,
		AuthIsRequired: false,
	},
	{
		URI:            "/users",
		Method:         http.MethodPost,
		Func:           controllers.CreateUser,
		AuthIsRequired: false,
	},
	{
		URI:            "/get-users",
		Method:         http.MethodGet,
		Func:           controllers.LoadUsersPage,
		AuthIsRequired: true,
	},
	{
		URI:            "/users/{userId}",
		Method:         http.MethodGet,
		Func:           controllers.LoadProfileUser,
		AuthIsRequired: true,
	},
	{
		URI:            "/users/{userId}/unfollow",
		Method:         http.MethodPost,
		Func:           controllers.UnfollowUser,
		AuthIsRequired: true,
	},
	{
		URI:            "/users/{userId}/follow",
		Method:         http.MethodPost,
		Func:           controllers.FollowUser,
		AuthIsRequired: true,
	},
}
