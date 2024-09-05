package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int
	Role      string
	Name      string
	Username  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
