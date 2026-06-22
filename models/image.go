package models

import (
	"time"

	"gorm.io/gorm"
)

type Image struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	UserID uint

	FileName string
	FilePath string

	User User
}
