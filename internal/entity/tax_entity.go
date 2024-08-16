package entity

import "time"

type Tax struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Rate      float64   `json:"rate" gorm:"type:decimal(5,2)"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (t *Tax) TableName() string {
	return "taxes"
}
