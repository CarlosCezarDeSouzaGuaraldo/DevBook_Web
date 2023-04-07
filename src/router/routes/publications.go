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
	{
		URI:            "/publications/{publicationId}/update",
		Method:         http.MethodGet,
		Func:           controllers.LoadEditPublicationScreen,
		AuthIsRequired: true,
	},
	{
		URI:            "/publications/{publicationId}",
		Method:         http.MethodPut,
		Func:           controllers.UpdatePublication,
		AuthIsRequired: true,
	},
	{
		URI:            "/publications/{publicationId}",
		Method:         http.MethodDelete,
		Func:           controllers.DeletePublication,
		AuthIsRequired: true,
	},
}
