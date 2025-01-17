package routes

import (
	"net/http"
	"ramada/api/src/auth"
	"ramada/api/src/controllers"
)

var userRoutes = []Route{
	{
		Uri:      "/login",
		Method:   http.MethodPost,
		Callback: auth.Login,
		authed:   false,
	},
	{
		Uri:      "/logout",
		Method:   http.MethodPost,
		Callback: auth.InvalidateToken,
		authed:   true,
	},
	{
		Uri:      "/users",
		Method:   http.MethodPost,
		Callback: controllers.UpinsertUser,
		authed:   false,
	},
	{
		Uri:      "/users/{id}",
		Method:   http.MethodPut,
		Callback: controllers.UpinsertUser,
		authed:   true,
	},
	{
		Uri:      "/me",
		Method:   http.MethodGet,
		Callback: controllers.GetCurrentUser,
		authed:   true,
	},
	{
		Uri:      "/users",
		Method:   http.MethodGet,
		Callback: controllers.ListUsers,
		authed:   true,
	},
	{
		Uri:      "/users/{id}",
		Method:   http.MethodGet,
		Callback: controllers.GetUser,
		authed:   true,
	},
	{
		Uri:      "/users/{id}",
		Method:   http.MethodDelete,
		Callback: controllers.DeleteUser,
		authed:   true,
	},
}
