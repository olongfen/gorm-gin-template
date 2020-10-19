package model

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	ID        uint            `json:"id" gorm:"primarykey"`
	CreatedAt *time.Time      `json:"createdAt,omitempty"`
	UpdatedAt *time.Time      `json:"updatedAt,omitempty"`
	DeletedAt *gorm.DeletedAt `json:"deletedAt,omitempty"gorm:"index"`
}
