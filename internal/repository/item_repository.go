package repository

import "waizly/internal/entity"

type ItemRepository struct {
	Repository[entity.Item]
}

func NewItemRepository() *ItemRepository {
	return &ItemRepository{}
}
