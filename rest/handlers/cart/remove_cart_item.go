package cart

import (
	"ecommerce/utils"
	"net/http"
	"strconv"
)

func (h *Handler) RemoveCartItem(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	itemID, _ := strconv.Atoi(idStr)
	userID := 1

	err := h.cartRepo.RemoveCartItem(userID, itemID)
	if err != nil {
		utils.SenError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	utils.SendData(w, map[string]string{"message": "Cart item removed"}, http.StatusOK)
}
