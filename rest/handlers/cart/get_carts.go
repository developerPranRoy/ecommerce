package cart

import (
	"ecommerce/utils"
	"net/http"
)

func (h *Handler) GetCart(w http.ResponseWriter, r *http.Request) {
	// user_id normally comes from JWT middleware; here assume 1
	userID := 1

	cartItems, err := h.cartRepo.GetCartItems(userID)
	if err != nil {
		utils.SenError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	utils.SendData(w, cartItems, http.StatusOK)
}
