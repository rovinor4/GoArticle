package model

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	Id        uint   `gorm:"primaryKey;autoIncrement"`
	Name      string `gorm:"size:255;not null"`
	Slug      string `gorm:"size:255;unique;not null"`
	Desc      string `gorm:"type:text"`
	UserId    uint
	Article   []Article      `gorm:"foreignKey:CategoryID;constraint:OnDelete:CASCADE"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
