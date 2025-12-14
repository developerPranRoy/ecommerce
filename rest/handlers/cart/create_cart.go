package cart

import (
	"ecommerce/repo"
	"ecommerce/utils"
	"encoding/json"
	"net/http"
)

type ReqAddToCart struct {
	ProductID int  `json:"product_id"`
	VariantID *int `json:"variant_id,omitempty"`
	Quantity  int  `json:"quantity"`
}

func (h *Handler) AddToCart(w http.ResponseWriter, r *http.Request) {
	var req ReqAddToCart
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Please provide valid JSON", http.StatusBadRequest)
		return
	}

	// user_id normally comes from JWT middleware; here assume 1
	userID := 1

	// TODO: Fetch actual product price from productRepo
	price := 499.0

	cartItem := repo.CartItem{
		UserID:       userID,
		ProductID:    req.ProductID,
		VariantID:    req.VariantID,
		Quantity:     req.Quantity,
		PriceAtAdded: price,
	}

	createdItem, err := h.cartRepo.CreateCart(cartItem)
	if err != nil {
		utils.SenError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	utils.SendData(w, createdItem, http.StatusCreated)
}
