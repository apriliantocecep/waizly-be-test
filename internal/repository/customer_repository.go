package repository

import "waizly/internal/entity"

type CustomerRepository struct {
	Repository[entity.Customer]
}

func NewCustomerRepository() *CustomerRepository {
	return &CustomerRepository{}
}
