package routes

import (
	"devbookapp/src/middlewares"
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
	routes = append(routes, homeRoutes...)
	routes = append(routes, publicationRoutes...)
	routes = append(routes, logoutRoutes...)

	for _, route := range routes {
		if route.RequiresAuthentication {
			r.HandleFunc(route.URI,
				middlewares.CheckAuthenticatedRequest(route.Function)).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI, route.Function).Methods(route.Method)
		}
	}

	fileServer := http.FileServer(http.Dir("./assets/"))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return r
}
