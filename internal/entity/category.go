package entity

import "time"

type Category struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Name        string `gorm:"unique;not null"` // Название темы
	Description string `gorm:"type:text"`       // Описание темы
}
