package cart

import (
	"ecommerce/utils"
	"encoding/json"
	"net/http"
	"strconv"
)

func (h *Handler) UpdateCartItem(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	itemID, _ := strconv.Atoi(idStr)

	var req struct {
		Quantity int `json:"quantity"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Please provide valid JSON", http.StatusBadRequest)
		return
	}
	userID := 1

	updatedItem, err := h.cartRepo.UpdateCartItem(userID, itemID, req.Quantity)
	if err != nil {
		utils.SenError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	utils.SendData(w, updatedItem, http.StatusOK)
}
