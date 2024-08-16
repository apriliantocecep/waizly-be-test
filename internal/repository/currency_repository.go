package repository

import "waizly/internal/entity"

type CurrencyRepository struct {
	Repository[entity.Currency]
}

func NewCurrencyRepository() *CurrencyRepository {
	return &CurrencyRepository{}
}
