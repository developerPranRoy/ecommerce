package cart

import (
	middlewares "ecommerce/rest/middleware"
	"net/http"
)

func (h *Handler) RegisterRoutes(
	mux *http.ServeMux,
	manager *middlewares.Manager) {

	mux.Handle("POST/cart/add",
		manager.With(
			http.HandlerFunc(h.AddToCart),
			h.middlewares.AuthenticationJWT,
		))

	mux.Handle("GET/cart",
		manager.With(
			http.HandlerFunc(h.GetCart),
			h.middlewares.AuthenticationJWT,
		))

	mux.Handle("PUT/cart/update",
		manager.With(
			http.HandlerFunc(h.UpdateCartItem),
			h.middlewares.AuthenticationJWT,
		))

	mux.Handle("DELETE/cart/remove",
		manager.With(
			http.HandlerFunc(h.RemoveCartItem),
			h.middlewares.AuthenticationJWT,
		))

	mux.Handle("DELETE/cart/clear",
		manager.With(
			http.HandlerFunc(h.ClearCart),
			h.middlewares.AuthenticationJWT,
		))

}
