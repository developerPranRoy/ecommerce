package products

import (
	"ecommerce/utils"
	"net/http"
	"strconv"
)

func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")
	id, err := strconv.Atoi(productID)
	if err != nil {
		utils.SenError(w, http.StatusBadRequest, " Plz Send Valid Id")
		return
	}
	product, err := h.productRepo.Get(id)
	if err != nil {
		utils.SenError(w, http.StatusInternalServerError, "Product not Found")
		return
	}
	if product == nil {
		utils.SenError(w, http.StatusNotFound, " Product Not Round ")
		return
	}
	utils.SendData(w, product, http.StatusOK)
}
