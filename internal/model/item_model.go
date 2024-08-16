package model

type CreateItemRequest struct {
	Name string `json:"name" binding:"required"`
	Type string `json:"type" binding:"required"`
}

type UpdateItemRequest struct {
	Name string `json:"name" binding:"required"`
	Type string `json:"type" binding:"required"`
}
