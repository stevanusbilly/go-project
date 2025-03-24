package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role struct {
	ID          uuid.UUID      `gorm:"type:uuid;primaryKey" json:"id"`
	Name        string         `gorm:"type:varchar(255);not null" json:"name"`
	Description string         `gorm:"type:varchar(255);not null;default:''" json:"description"`
	CreatedAt   time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	CreatedBy   uuid.UUID      `gorm:"type:uuid" json:"created_by"`
	UpdatedAt   time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	UpdatedBy   uuid.UUID      `gorm:"type:uuid" json:"updated_by"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	DeletedBy   uuid.UUID      `gorm:"type:uuid" json:"deleted_by"`
}

func (Role) TableName() string {
	return "role"
}