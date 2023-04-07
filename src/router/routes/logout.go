package routes

import (
	"net/http"
	"web/src/controllers"
)

var logoutRoute = Route{
	URI:            "/logout",
	Method:         http.MethodGet,
	Func:           controllers.DoLogout,
	AuthIsRequired: false,
}
