package entity

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	ID          int
	UserID      int
	ProjectID   int
	Title       string
	Description string
	Status      string
	Priority    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}
