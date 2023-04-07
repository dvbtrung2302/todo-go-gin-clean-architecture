package entity

import "time"

type Todo struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Status    string    `json:"status" gorm:"default:new"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
