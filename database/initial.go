package database

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id uint `gorm:"primaryKey" json:"id" validate:"omitempty,number"`

	Firstname string `gorm:"not null; type:varchar(255)" json:"firstname" validate:"required,max=255"`
	Lastname  string `gorm:"not null; type:varchar(255)" json:"lastname" validate:"required,max=255"`
	Email     string

	//Timestamps
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`

	//Many 2 many FK
	Roles []Role `gorm:"many2many:user_roles;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Role struct {
	Id uint `gorm:"primaryKey" json:"id" validate:"omitempty,number"`

	Label     string    `gorm:"not null; type:varchar(255)" json:"label" validate:"required,max=255"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	//Many 2 many FK
	Users []User `gorm:"many2many:user_roles;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
