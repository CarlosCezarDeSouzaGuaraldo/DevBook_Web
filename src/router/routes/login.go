package routes

import (
	"net/http"
	"web/src/controllers"
)

var loginRoutes = []Route{
	{
		URI:            "/",
		Method:         http.MethodGet,
		Func:           controllers.LoadLoginScreen,
		AuthIsRequired: false,
	},
	{
		URI:            "/login",
		Method:         http.MethodGet,
		Func:           controllers.LoadLoginScreen,
		AuthIsRequired: false,
	},
	{
		URI:            "/login",
		Method:         http.MethodPost,
		Func:           controllers.DoLogin,
		AuthIsRequired: false,
	},

}
