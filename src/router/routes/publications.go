package routes

import (
	"net/http"
	"web/src/controllers"
)

var publicationRoutes = []Route{
	{
		URI:            "/publications",
		Method:         http.MethodPost,
		Func:           controllers.CreatePublication,
		AuthIsRequired: true,
	},
}
