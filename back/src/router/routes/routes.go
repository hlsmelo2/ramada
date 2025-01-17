package routes

import (
	"net/http"
	"ramada/api/src/middlewares"
	"slices"

	"github.com/gorilla/mux"
)

type Route struct {
	Uri      string
	Method   string
	Callback func(http.ResponseWriter, *http.Request)
	authed   bool
}

func Setup(router *mux.Router) *mux.Router {
	routes := slices.Concat(userRoutes, productRoutes)

	for _, route := range routes {
		if route.authed {
			router.HandleFunc(route.Uri, middlewares.Authenticate(route.Callback)).Methods(route.Method)

			continue
		}

		router.HandleFunc(route.Uri, route.Callback).Methods(route.Method)
	}

	return router
}
