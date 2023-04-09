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
	{
		URI:            "/profile",
		Method:         http.MethodGet,
		Func:           controllers.LoadLoggedProfileUser,
		AuthIsRequired: true,
	},
	{
		URI:            "/edit-user",
		Method:         http.MethodGet,
		Func:           controllers.LoadEditProfileUserScreen,
		AuthIsRequired: true,
	},
	{
		URI:            "/edit-user",
		Method:         http.MethodPut,
		Func:           controllers.EditUser,
		AuthIsRequired: true,
	},
	{
		URI:            "/update-password",
		Method:         http.MethodGet,
		Func:           controllers.LoadUpdatePasswordScreen,
		AuthIsRequired: true,
	},
	{
		URI:            "/update-password",
		Method:         http.MethodPost,
		Func:           controllers.UpdatePassword,
		AuthIsRequired: true,
	},
}
