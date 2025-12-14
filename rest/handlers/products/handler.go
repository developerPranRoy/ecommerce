package products

import (
	"ecommerce/repo"
	middlewares "ecommerce/rest/middleware"
)

type Handler struct {
	middleware  middlewares.Middlewares
	productRepo repo.ProductRepo
	
}

func NewHandler(middlewares *middlewares.Middlewares,
	productRepo repo.ProductRepo) *Handler {
	return &Handler{
		middleware: *middlewares,
		productRepo: productRepo,

	}
}
