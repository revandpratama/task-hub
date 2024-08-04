package entity

import "time"

type Project struct {
	ID          int
	UserID      int
	Title       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}
