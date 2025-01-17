package router

import (
	"ramada/api/src/router/routes"

	"github.com/gorilla/mux"
)

// return a initialized router instance
func GetRouter() *mux.Router {
	router := mux.NewRouter()
	router.StrictSlash(true)

	return routes.Setup(router)
}
