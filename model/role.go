package model

import (
	"time"
)

type Role struct {
	Id uint `gorm:"primaryKey" json:"id" validate:"omitempty,number"`

	Label     string    `gorm:"not null; type:varchar(255)" json:"label" validate:"required,max=255"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	//Has many FK
	Users []User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
