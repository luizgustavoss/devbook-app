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
	{
		URI:                    "/search-users",
		Method:                 http.MethodGet,
		Function:               controllers.SearchUsers,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/users/{userID}",
		Method:                 http.MethodGet,
		Function:               controllers.GetUserDetails,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/users/{userID}/follow",
		Method:                 http.MethodPost,
		Function:               controllers.FollowUser,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/users/{userID}/unfollow",
		Method:                 http.MethodPost,
		Function:               controllers.UnfollowUser,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/profile",
		Method:                 http.MethodGet,
		Function:               controllers.LoadLoggedUserProfile,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/edit-user",
		Method:                 http.MethodGet,
		Function:               controllers.LoadEditUserPage,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/change-password",
		Method:                 http.MethodGet,
		Function:               controllers.LoadChangePasswordPage,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/users",
		Method:                 http.MethodPut,
		Function:               controllers.UpdateUser,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/change-password",
		Method:                 http.MethodPost,
		Function:               controllers.ChangePassword,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/delete-account",
		Method:                 http.MethodDelete,
		Function:               controllers.DeleteAccount,
		RequiresAuthentication: true,
	},

}