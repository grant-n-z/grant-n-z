package entity

import "time"

type ItemEntity struct {
	ID        uint64     `gorm:"primary_key" json:"id"`
	Name      string     `json:"name"`
	Category  string     `json:"category"`
	CreatedAt time.Time  `json:"created_at"`
}
