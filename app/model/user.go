package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id                uint           `gorm:"primaryKey;autoIncrement"`
	Username          string         `gorm:"size:255;unique;not null"`
	Email             string         `gorm:"size:255;unique;not null"`
	Password          string         `gorm:"not null"`
	DisplayName       string         `gorm:"size:255"`
	Bio               string         `gorm:"type:text"`
	ProfilePictureUrl string         `gorm:"size:255"`
	Role              string         `gorm:"type:enum('admin','user');default:'user'"`
	Article           []Article      `gorm:"foreignKey:AuthorID;constraint:OnDelete:CASCADE"`
	Token             []Token        `gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE"`
	CreatedAt         time.Time      `gorm:"autoCreateTime"`
	UpdatedAt         time.Time      `gorm:"autoUpdateTime"`
	DeleteAt          gorm.DeletedAt `gorm:"index"`
}

type UserRegister struct {
	Username          string `json:"username" validate:"required,min=3,uniqueUsername" `
	Email             string `json:"email" validate:"required,email,uniqueEmail"`
	Password          string `json:"password" validate:"required,min=8"`
	DisplayName       string `json:"display_name" validate:"required,min=3"`
	Bio               string `json:"bio" validate:"required,min=3"`
	ProfilePictureUrl string `json:"profile_picture_url" validate:"required,url"`
}

type UserLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
