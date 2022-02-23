package entity

import (
	"gorm.io/gorm"
	"time"
)

type Book struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Title     string         `json:"title"`
	Author    string         `json:"author"`
	Cover     string         `json:"cover"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index,column:deleted_at"`
}
