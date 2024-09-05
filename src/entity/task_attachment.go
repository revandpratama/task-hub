package entity

import (
	"time"

	"gorm.io/gorm"
)

type TaskAttachment struct {
	TaskID    int
	UserID    int
	FilePath  string
	FileType  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
