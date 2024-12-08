package entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	TelegramID string `gorm:"unique;not null"` // Уникальный идентификатор Telegram
	Username   string `gorm:"not null"`        // Имя пользователя
}
