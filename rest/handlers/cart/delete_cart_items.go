package cart

import (
	"ecommerce/utils"
	"net/http"
)

func (h *Handler) ClearCart(w http.ResponseWriter, r *http.Request) {
	userID := 1
	err := h.cartRepo.ClearCart(userID)
	if err != nil {
		utils.SenError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	utils.SendData(w, map[string]string{"message": "Cart cleared"}, http.StatusOK)
}
