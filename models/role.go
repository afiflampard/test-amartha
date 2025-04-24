package models

import (
	"time"

	"github.com/google/uuid"
)

type Role struct {
	ID        uuid.UUID `gorm:"column:id" json:"id"`
	Name      string    `gorm:"column:name" json:"name"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (Role) TableName() string {
	return "role"
}
