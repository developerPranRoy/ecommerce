package users

import (
	middlewares "ecommerce/rest/middleware"
	"net/http"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middlewares.Manager) {

	mux.Handle("POST /users/register",
		manager.With(http.HandlerFunc(h.CreateUser)))

	mux.Handle("POST /users/login",
		manager.With(http.HandlerFunc(h.Login)))
}
