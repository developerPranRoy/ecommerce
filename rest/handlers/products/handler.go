package products

import (
	middlewares "ecommerce/rest/middleware"
)

type Handler struct {
	middleware middlewares.Middlewares
	svc        Service
}

func NewHandler(middlewares *middlewares.Middlewares,
	svc Service) *Handler {
	return &Handler{
		middleware: *middlewares,
		svc:        svc,
	}
}
