package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID                 string    `gorm:"type:uuid;primary_key;uuid_generate_v4;" json:"id"  form:"id"`
	FullName           string    `gorm:"type:varchar;not null"  json:"full_name" form:"full_name" validate:"required"`
	ProfileImage       string    `gorm:"type:text" json:"profile_image" form:"profile_image"`
	Username           string    `gorm:"type:varchar(100);not null;unique" json:"username" form:"username" validate:"required,max=100"`
	Password           string    `gorm:"type:text;not null" json:"password" form:"password" validate:"required"`
	TwoFactorToken     string    `json:"two_factor_token" form:"two_factor_token"`
	IsTwoFactorEnabled bool      `gorm:"default:false" json:"is_two_factor_enabled" form:"is_two_factor_enabled"`
	InvalidSignInCount int       `gorm:"default:0" json:"invalid_signIn_count" form:"invalid_signIn_count"`
	Status             bool      `gorm:"default:false" json:"status" form:"status"` // Status can be "active" or "locked"
	CreatedAt          time.Time `json:"created_at" form:"created_at"`
	UpdatedAt          time.Time `json:"updated_at" form:"updated_at"`
}

func (item *User) BeforeCreate(tx *gorm.DB) (err error) {
	item.ID = uuid.New().String()
	return nil
}
