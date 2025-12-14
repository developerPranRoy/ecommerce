package reviews

import (
	"ecommerce/utils"
	"net/http"
)

func (h *Handler) GetReviews(w http.ResponseWriter, r *http.Request) {
	reviewsList, err := h.reviewRepo.ListAll()
	if err != nil {
		utils.SenError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	utils.SendData(w, reviewsList, http.StatusOK)
}
