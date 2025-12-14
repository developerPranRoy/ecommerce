package reviews

import "ecommerce/repo"

type Handler struct {
	reviewRepo repo.ReviewRepo
}

func NewHandler(reviewRepo repo.ReviewRepo) *Handler {
	return &Handler{
		reviewRepo: reviewRepo,
	}
}
