package model

import (
	"time"

	"gorm.io/gorm"
)

type Article struct {
	Id         uint   `gorm:"primaryKey;autoIncrement"`
	Title      string `gorm:"size:255;unique;not null"`
	Context    string `gorm:"type:text"`
	ViewCount  int    `gorm:"default:0"`
	Slug       string `gorm:"size:255;unique;not null"`
	AuthorID   uint
	CategoryID uint
	CreatedAt  time.Time      `gorm:"autoCreateTime"`
	UpdatedAt  time.Time      `gorm:"autoUpdateTime"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
