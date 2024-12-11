package entity

import "time"

type Category struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time

	UserID      uint   `gorm:"not null"`        // Внешний ключ на пользователя
	Name        string `gorm:"unique;not null"` // Название темы
	Description string `gorm:"type:text"`       // Описание темы

	User User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"` // Связь с пользователем
}
