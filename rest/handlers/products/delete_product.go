package products

import (
	"ecommerce/utils"
	"net/http"
	"strconv"
)

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")
	id, err := strconv.Atoi(productID)
	if err != nil {
		utils.SenError(w, http.StatusBadRequest, " Plz Send Valid Id")
		return
	}

	err = h.svc.Delete(id)
	if err != nil {
		utils.SenError(w, http.StatusInternalServerError, "Internal Server Error")
		return

	}

	// database.Delete(id)
	utils.SendData(w, " Data Deleted", http.StatusAccepted)

}
