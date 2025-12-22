package products

import (
	"ecommerce/utils"
	"fmt"
	"net/http"
	"strconv"
)

var cnt int

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {

	rqQuery := r.URL.Query()
	pageString := rqQuery.Get("page")
	limitString := rqQuery.Get("page")
	page, _ := strconv.ParseInt(pageString, 10, 32)
	limit, _ := strconv.ParseInt(limitString, 10, 32)
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}

	productList, err := h.svc.List(int(page), int(limit))
	if err != nil {
		utils.SenError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	go func() {
		count, err := h.svc.Count()
		if err != nil {
			utils.SenError(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		cnt = count
	}()
	go func() {
		count1, err := h.svc.Count()
		if err != nil {
			utils.SenError(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		fmt.Println(count1)
	}()

	go func() {
		count2, err := h.svc.Count()
		if err != nil {
			utils.SenError(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		fmt.Println(count2)
	}()

	// time.Sleep(2 * time.Second)
	utils.SendPage(w, productList, int(page), int(limit), cnt)
}
