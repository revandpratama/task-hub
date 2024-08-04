package entity

import "time"

type Task struct {
	ID          int
	Title       string
	Description string
	Status      string
	Priority    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}
