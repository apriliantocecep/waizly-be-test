package model

import "waizly/internal/entity"

type CurrencyResponse struct {
	Currency *entity.Currency `json:"currency"`
}

type CreateCurrencyRequest struct {
	Code         string  `json:"code" binding:"required"`
	Name         string  `json:"name" binding:"required"`
	ExchangeRate float64 `json:"exchange_rate" binding:"required,gt=0"`
}

type UpdateCurrencyRequest struct {
	Code         string  `json:"code" binding:"required"`
	Name         string  `json:"name" binding:"required"`
	ExchangeRate float64 `json:"exchange_rate" binding:"required,gt=0"`
}
