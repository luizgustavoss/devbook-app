package routes

import (
	"devbookapp/src/controllers"
	"net/http"
)

var homeRoutes = []Route {
	{
		URI: "/home",
		Method: http.MethodGet,
		Function: controllers.LoadHomePage,
		RequiresAuthentication: true,
	},
}