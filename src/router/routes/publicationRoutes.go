package routes

import (
	"devbookapp/src/controllers"
	"net/http"
)

var publicationRoutes = []Route {
	{
		URI: "/publications",
		Method: http.MethodPost,
		Function: controllers.CreatePublication,
		RequiresAuthentication: true,
	},
	{
		URI: "/publications/{publicationId}/like",
		Method: http.MethodPost,
		Function: controllers.LikePublication,
		RequiresAuthentication: true,
	},
	{
		URI: "/publications/{publicationId}/unlike",
		Method: http.MethodPost,
		Function: controllers.UnlikePublication,
		RequiresAuthentication: true,
	},
	{
		URI: "/publications/{publicationId}/edit",
		Method: http.MethodGet,
		Function: controllers.LoadPublicationDetail,
		RequiresAuthentication: true,
	},
	{
		URI: "/publications/{publicationId}",
		Method: http.MethodDelete,
		Function: controllers.DeletePublication,
		RequiresAuthentication: true,
	},
	{
		URI: "/publications/{publicationId}",
		Method: http.MethodPut,
		Function: controllers.UpdatePublication,
		RequiresAuthentication: true,
	},

}
