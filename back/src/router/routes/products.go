package routes

import (
	"net/http"
	"ramada/api/src/controllers"
)

var productRoutes = []Route{
	{
		Uri:      "/products",
		Method:   http.MethodGet,
		Callback: controllers.GetProducts,
		authed:   true,
		// ?nome=...categoria=...preco_min=...preco_max=...
	},
	{
		Uri:      "/products/{id}",
		Method:   http.MethodGet,
		Callback: controllers.GetProduct,
		authed:   true,
	},
	{
		Uri:      "/products",
		Method:   http.MethodPost,
		Callback: controllers.UpinsertProduct,
		authed:   true,
	},
	{
		Uri:      "/products/importar/{id}",
		Method:   http.MethodPost,
		Callback: controllers.ImportProducts,
		authed:   true,
	},
	{
		Uri:      "/products/{id}",
		Method:   http.MethodPut,
		Callback: controllers.UpinsertProduct,
		authed:   true,
	},
	{
		Uri:      "/products/{id}",
		Method:   http.MethodDelete,
		Callback: controllers.DeleteProduct,
		authed:   true,
	},
}
