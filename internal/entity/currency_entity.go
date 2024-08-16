package entity

import "time"

type Currency struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	Code         string    `json:"code"`
	Name         string    `json:"name"`
	ExchangeRate float64   `json:"exchange_rate" gorm:"type:decimal(10,2)"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (t *Currency) TableName() string {
	return "currencies"
}
