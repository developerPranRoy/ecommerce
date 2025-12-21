package utils

import (
	"net/http"
)

type PaginatedData struct {
	Data       any        `json:"data"`
	Pagination Pagination `json:"pagination"`
}

type Pagination struct {
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	TotalItems int `json:"totalItems"`
	TotalPage  int `json:"totalPage"`
}

func SendPage(w http.ResponseWriter, data any, page, limit, count int) {

	paginatedData := PaginatedData{
		Data: data,
		Pagination: Pagination{
			Page:       int(page),
			Limit:      int(limit),
			TotalItems: count,
			TotalPage:  count / int(limit),
		},
	}
	SendData(w, paginatedData, http.StatusOK)
}
