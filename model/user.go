package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id uint `gorm:"primaryKey" json:"id" validate:"omitempty"`

	//FK
	RoleId uint `gorm:"not null" json:"role_id" validate:"omitempty"`
	Role   Role `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	Firstname string `gorm:"not null; type:varchar(255)" json:"firstname" validate:"required,max=255"`
	Lastname  string `gorm:"not null; type:varchar(255)" json:"lastname" validate:"required,max=255"`
	Email     string `gorm:"not null; type:varbinary(255)" json:"email" validate:"required,max=255"`
	Password  string `gorm:"not null; type:varchar(255)" json:"password" validate:"required,max=255"`

	//Timestamps
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
