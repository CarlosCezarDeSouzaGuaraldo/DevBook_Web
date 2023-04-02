package routes

import (
	"net/http"

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

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Func).Methods(route.Method)
	}

	fileServer := http.FileServer(http.Dir("./assets/"))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return r
}
