package products

import (
	middlewares "ecommerce/rest/middleware"
	"net/http"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middlewares.Manager) {

	mux.Handle("GET /products",
		manager.With(http.HandlerFunc(h.GetProducts)))

	mux.Handle("POST /create_products",
		manager.With(http.HandlerFunc(h.CreateProducts),
			h.middleware.AuthenticationJWT,
		))

	mux.Handle("GET /products/{id}",
		manager.With(http.HandlerFunc(h.GetProduct)))

	mux.Handle("PUT /products/{id}",
		manager.With(http.HandlerFunc(h.UpdateProduct),
			h.middleware.AuthenticationJWT,
		))

	mux.Handle("DELETE /products/{id}",
		manager.With(http.HandlerFunc(h.UpdateProduct),
			h.middleware.AuthenticationJWT,
		))

}
