package entity

import (
	"gorm.io/gorm"
)

type Question struct {
	gorm.Model
	CategoryID uint   `gorm:"not null"`           // Внешний ключ на категорию
	Text       string `gorm:"type:text;not null"` // Текст вопроса
	AnswerJSON string `gorm:"type:json;not null"` // JSON с правильным ответом

	Category Category `gorm:"foreignKey:CategoryID;constraint:OnDelete:CASCADE"` // Связь с категорией
}
