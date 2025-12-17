package products

import (
	"ecommerce/utils"
	"net/http"
	"strconv"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	rqQuery := r.URL.Query()
	pageString := rqQuery.Get("page")
	limitString := rqQuery.Get("page")
	page, _ := strconv.ParseInt(pageString, 10, 64)
	limit, _ := strconv.ParseInt(limitString, 10, 64)
	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 10
	}

	productList, err := h.svc.List(int(page), int(limit))
	if err != nil {
		utils.SenError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	utils.SendData(w, productList, http.StatusOK)
}
