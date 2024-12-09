package entity

import (
	"gorm.io/gorm"
)

type Test struct {
	gorm.Model
	UserID    uint   `gorm:"not null"`           // Внешний ключ на пользователя
	Questions string `gorm:"type:json;not null"` // JSON с массивом ID вопросов
	Score     int    `gorm:"default:0"`          // Счет (правильные ответы)

	User User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"` // Связь с пользователем
}
