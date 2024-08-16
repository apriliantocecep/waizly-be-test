package model

type CreateCustomerRequest struct {
	Name    string `json:"name" binding:"required"`
	Status  string `json:"status" binding:"required"`
	Address string `json:"address" binding:"required"`
}

type UpdateCustomerRequest struct {
	Name    string `json:"name" binding:"required"`
	Status  string `json:"status" binding:"required"`
	Address string `json:"address" binding:"required"`
}
