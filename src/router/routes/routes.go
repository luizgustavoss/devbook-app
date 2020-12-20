package routes

import (
	"github.com/gorilla/mux"
	"net/http"
)

// represents WEB APP routes
type Route struct {
	URI                    string
	Method                 string
	Function               func(http.ResponseWriter, *http.Request)
	RequiresAuthentication bool
}

// configures routes in router
func ConfigureRoutes(r *mux.Router) *mux.Router {

	routes := loginRoutes
	routes = append(routes, userRoutes...)

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	fileServer := http.FileServer(http.Dir("./assets/"))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return r
}
