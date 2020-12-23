package routes

import (
	"devbookapp/src/controllers"
	"net/http"
)

var logoutRoutes = []Route {
	{
		URI: "/logout",
		Method: http.MethodGet,
		Function: controllers.Logout,
		RequiresAuthentication: true,
	},
}