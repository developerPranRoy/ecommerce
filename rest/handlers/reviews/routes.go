package reviews

import (
	middlewares "ecommerce/rest/middleware"
	"net/http"
)

func (h Handler) RegisterRoutes(mux *http.ServeMux, manager *middlewares.Manager) {

	mux.Handle("GET /reviews",
		manager.With(http.HandlerFunc(h.GetReviews)))

	mux.Handle("POST /create_reviews",
		manager.With(http.HandlerFunc(h.CreateReview)))

	// mux.Handle("GET /update_reviews",
	// 	manager.With(http.HandlerFunc(h.GetReview)))
}
