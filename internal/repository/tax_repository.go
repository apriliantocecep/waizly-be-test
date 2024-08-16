package repository

import "waizly/internal/entity"

type TaxRepository struct {
	Repository[entity.Tax]
}

func NewTaxRepository() *TaxRepository {
	return &TaxRepository{}
}
