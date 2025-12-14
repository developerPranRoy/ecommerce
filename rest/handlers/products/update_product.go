package products

import (
	"ecommerce/repo"
	"ecommerce/utils"
	"encoding/json"
	"net/http"
	"strconv"
)

type ReqUpdateProduct struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImageUrl    string  `json:"imageUrl"`
}

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")
	id, err := strconv.Atoi(productID)
	if err != nil {
		utils.SenError(w, http.StatusBadRequest, " Plz Send Valid Id")
		return
	}

	var req ReqCreateProduct

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&req)
	if err != nil {
		utils.SenError(w, http.StatusBadRequest, "Send Valid Json")
		return
	}
	_, err = h.productRepo.Update(repo.Product{
		ID:          id,
		Title:       req.Title,
		Description: req.Description,
		Price:       req.Price,
		ImageUrl:    req.ImageUrl,
	})
	if err != nil {
		utils.SenError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	utils.SendData(w, " Data Updated", http.StatusAccepted)
}
