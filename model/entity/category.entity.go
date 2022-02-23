package entity

import (
	"gorm.io/gorm"
	"time"
)

type Category struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name"`
	Photos    []Photo        `json:"photos"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index,column:deleted_at"`
}
