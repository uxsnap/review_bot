package entity

import (
	"gorm.io/gorm"
)

type UsefulLink struct {
	gorm.Model
	URL         string `gorm:"not null"`  // Ссылка
	Description string `gorm:"type:text"` // Описание ссылки
}
