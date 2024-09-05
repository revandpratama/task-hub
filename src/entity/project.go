package entity

import (
	"time"

	"gorm.io/gorm"
)

type Project struct {
	ID          int
	UserID      int
	Title       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}
