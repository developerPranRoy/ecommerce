package cart

import (
	"ecommerce/repo"
	middlewares "ecommerce/rest/middleware"
)

type Handler struct {
	cartRepo    repo.CartRepo
	middlewares middlewares.Middlewares
}

func NewCartHandler(
	cartRepo repo.CartRepo,
	middlewares *middlewares.Middlewares,
) *Handler {
	return &Handler{
		cartRepo:    cartRepo,
		middlewares: *middlewares,
	}
}
