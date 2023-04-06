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
	{
		URI:            "/publications/{publicationId}/like",
		Method:         http.MethodPost,
		Func:           controllers.LikePublication,
		AuthIsRequired: true,
	},
	{
		URI:            "/publications/{publicationId}/unlike",
		Method:         http.MethodPost,
		Func:           controllers.UnlikePublication,
		AuthIsRequired: true,
	},
}
