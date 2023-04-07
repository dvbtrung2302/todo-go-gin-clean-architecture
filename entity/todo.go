package entity

import "time"

type Todo struct {
	ID        string
	Status    string
	Content   string
	CreatedAt time.Time
	UpdateAt  time.Time
}
