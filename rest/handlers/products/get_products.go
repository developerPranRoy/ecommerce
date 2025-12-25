package products

import (
	"ecommerce/utils"
	"net/http"
	"strconv"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {

	rqQuery := r.URL.Query()
	pageString := rqQuery.Get("page")
	limitString := rqQuery.Get("limit")
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

	ch := make(chan int)
	// var wg sync.WaitGroup

	go func() {
		count, err := h.svc.Count()
		if err != nil {
			utils.SenError(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		ch <- count

	}()
	totalCount := <-ch

	// time.Sleep(2 * time.Second)
	// wg.Wait()
	utils.SendPage(w, productList, int(page), int(limit), totalCount)
}
