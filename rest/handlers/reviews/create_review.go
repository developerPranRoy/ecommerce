package reviews

import (
	"ecommerce/repo"
	"ecommerce/utils"
	"encoding/json"
	"net/http"
)

type ReqReview struct {
	ID        int    `json:"id"`
	ProductID int    `json:"product_id"`
	UserID    int    `json:"user_id"`
	Rating    int    `json:"rating"` // 1–5
	Comment   string `json:"comment"`
	// Images    []string `json:"images,omitempty"`
}

func (h Handler) CreateReview(w http.ResponseWriter, r *http.Request) {
	var req ReqReview

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		utils.SenError(w, http.StatusBadRequest, "Error ")
		return
	}
	createdReview, err := h.reviewRepo.Create(repo.Review{
		ProductID: req.ProductID,
		UserID:    req.UserID,
		Rating:    req.Rating,
		Comment:   req.Comment,
	})
	if err != nil {
		utils.SenError(w, http.StatusInternalServerError, "Internal Server Error")
		return

	}
	utils.SendData(w, createdReview, http.StatusOK)
}
