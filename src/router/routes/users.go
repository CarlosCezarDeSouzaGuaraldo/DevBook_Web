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
}
