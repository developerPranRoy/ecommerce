package products

import (
	"ecommerce/repo"
	"ecommerce/utils"
	"encoding/json"
	"net/http"
)

type ReqCreateProduct struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImageUrl    string  `json:"imageUrl"`
}

func (h *Handler) CreateProducts(w http.ResponseWriter, r *http.Request) {
	var reqProduct ReqCreateProduct
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&reqProduct)
	if err != nil {
		http.Error(w, "Plz give valid json", http.StatusBadRequest)
		return
	}
	createrdProduct, err := h.productRepo.Create(repo.Product{
		Title:       reqProduct.Title,
		Description: reqProduct.Description,
		Price:       reqProduct.Price,
		ImageUrl:    reqProduct.ImageUrl,
	})
	if err != nil {
		utils.SenError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	utils.SendData(w, createrdProduct, http.StatusCreated)

}
