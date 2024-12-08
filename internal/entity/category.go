package entity

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name        string `gorm:"unique;not null"` // Название темы
	Description string `gorm:"type:text"`       // Описание темы
}
