package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Task struct {
	ID          string    `gorm:"type:uuid;primary_key;uuid_generate_v4;" json:"id"  form:"id"`
	Title       string    `gorm:"type:varchar;not null"  json:"title" form:"title" validate:"required"`
	Description string    `gorm:"type:text" json:"description" form:"description"`
	Status      string    `gorm:"default:'in_progress'" json:"status" form:"status"`
	CreatedAt   time.Time `json:"created_at" form:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" form:"updated_at"`
}

func (task *Task) BeforeCreate(tx *gorm.DB) (err error) {
	task.ID = uuid.New().String()
	return nil
}
