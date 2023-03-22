package entity

import "time"

type Todo struct {
	Id        string
	Status    string
	Content   string
	CreatedAt time.Time
	UpdateAt  time.Time
}
