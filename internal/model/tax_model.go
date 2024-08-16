package model

import "waizly/internal/entity"

type TaxResponse struct {
	Tax *entity.Tax `json:"tax"`
}

type CreateTaxRequest struct {
	Name string  `json:"name" binding:"required"`
	Rate float64 `json:"rate" binding:"required,gt=0"`
}

type UpdateTaxRequest struct {
	Name string  `json:"name" binding:"required"`
	Rate float64 `json:"rate" binding:"required,gt=0"`
}
