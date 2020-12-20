package routes

import (
	"devbookapp/src/controllers"
	"net/http"
)

var userRoutes = []Route{
	{
		URI:                    "/create-user",
		Method:                 http.MethodGet,
		Function:               controllers.LoadCreateUserPage,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/users",
		Method:                 http.MethodPost,
		Function:               controllers.CreateUser,
		RequiresAuthentication: false,
	},
}