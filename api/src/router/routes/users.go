package routes

import (
	"api/src/controller"
	"net/http"
)

var routesUsers = []Route{
	{
		URI:                    "/users",
		Method:                 http.MethodPost,
		Function:               controller.CreateUser,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/users",
		Method:                 http.MethodGet,
		Function:               controller.FindAllUsers,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/users/{id}",
		Method:                 http.MethodGet,
		Function:               controller.FindUserById,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/users/{id}",
		Method:                 http.MethodPut,
		Function:               controller.UpdateUser,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/users/{id}",
		Method:                 http.MethodDelete,
		Function:               controller.DeleteUser,
		RequiresAuthentication: false,
	},
}
