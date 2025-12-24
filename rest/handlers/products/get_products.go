package products

import (
	"ecommerce/utils"
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

var cnt int
var mu sync.Mutex

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

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		mu.Lock()
		defer wg.Done()
		defer mu.Unlock()

		count, err := h.svc.Count()
		if err != nil {
			utils.SenError(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		cnt = count

	}()
	wg.Add(1)
	go func() {
		mu.Lock()
		defer wg.Done()
		defer mu.Unlock()

		count1, err := h.svc.Count()
		if err != nil {
			utils.SenError(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		fmt.Println(count1)

	}()

	wg.Add(1)
	go func() {
		mu.Lock()
		defer wg.Done()
		defer mu.Unlock()

		count2, err := h.svc.Count()
		if err != nil {
			utils.SenError(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		fmt.Println(count2)

	}()

	// time.Sleep(2 * time.Second)
	wg.Wait()
	utils.SendPage(w, productList, int(page), int(limit), cnt)
}
