package routes

import (
	"net/http"
	"web/src/middlewares"

	"github.com/gorilla/mux"
)

// Route represents all API's routes
type Route struct {
	URI            string
	Method         string
	Func           func(http.ResponseWriter, *http.Request)
	AuthIsRequired bool
}

// Configure all routes on application
func Configure(r *mux.Router) *mux.Router {
	routes := loginRoutes
	routes = append(routes, routeHomePage)
	routes = append(routes, userRoutes...)
	routes = append(routes, publicationRoutes...)
	routes = append(routes, logoutRoute)

	for _, route := range routes {
		if route.AuthIsRequired {
			r.HandleFunc(
				route.URI, middlewares.Logger(middlewares.Auth(route.Func)),
			).Methods(route.Method)
		} else {
			r.HandleFunc(
				route.URI, middlewares.Logger(route.Func),
			).Methods(route.Method)
		}
	}

	fileServer := http.FileServer(http.Dir("./assets/"))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return r
}
