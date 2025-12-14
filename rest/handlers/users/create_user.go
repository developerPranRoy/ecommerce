package users

import (
	"ecommerce/domain"
	"ecommerce/utils"
	"encoding/json"
	"net/http"
)

type ReqCreateUser struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Isowner   bool   `json:"isowner"`
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req ReqCreateUser
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		utils.SenError(w, http.StatusBadRequest, "Invalid Request Body")
		return
	}
	usr, err := h.svc.Create(domain.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  req.Password,
		Isowner:   req.Isowner,
	})
	if err != nil {
		utils.SenError(w, http.StatusInternalServerError, "Internal server error")
		return
	}
	utils.SendData(w, usr, http.StatusCreated)
}
