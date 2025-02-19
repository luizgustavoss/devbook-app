package routes

import (
	"devbookapp/src/controllers"
	"net/http"
)

var loginRoutes = []Route {
	{
		URI: "/",
		Method: http.MethodGet,
		Function: controllers.LoadLoginPage,
		RequiresAuthentication: false,
	},
	{
		URI: "/login",
		Method: http.MethodGet,
		Function: controllers.LoadLoginPage,
		RequiresAuthentication: false,
	},
	{
		URI: "/login",
		Method: http.MethodPost,
		Function: controllers.Login,
		RequiresAuthentication: false,
	},
}