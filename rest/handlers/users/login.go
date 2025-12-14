package users

import (
	"ecommerce/utils"
	"encoding/json"
	"net/http"
)

type ReqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var req ReqLogin
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		utils.SenError(w, http.StatusBadRequest, "Invalid credintials")
		return
	}
	usr, err := h.svc.Find(req.Email, req.Password)
	if err != nil {
		utils.SenError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}
	if usr == nil {
		utils.SenError(w, http.StatusBadRequest, "Invalid data")
		return
	}
	accesToken, err := utils.CreateJWT(h.cnf.JWTSecret, utils.Payload{
		Sub:       usr.ID,
		FirstName: usr.FirstName,
		LastName:  usr.LastName,
		Email:     usr.Email,
	})
	if err != nil {
		utils.SenError(w, http.StatusBadRequest, "Internals server error")
		return
	}
	utils.SendData(w, accesToken, http.StatusCreated)
}
