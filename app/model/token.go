package model

import (
	"time"

	"gorm.io/gorm"
)

type Token struct {
	Id        uint   `gorm:"primaryKey;autoIncrement"`
	Token     string `gorm:"size:255;not null"`
	UserId    uint
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
