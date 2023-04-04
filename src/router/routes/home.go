package routes

import (
	"net/http"
	"web/src/controllers"
)

var routeHomePage = Route{
	URI:    "/home",
	Method: http.MethodGet,
	Func: controllers.LoadHomePage,
	AuthIsRequired: true,
}