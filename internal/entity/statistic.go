package entity

import (
	"gorm.io/gorm"
)

type Statistic struct {
	gorm.Model
	UserID           uint `gorm:"not null"`  // Внешний ключ на пользователя
	TotalQuestions   int  `gorm:"default:0"` // Всего решено вопросов
	TotalTestsPassed int  `gorm:"default:0"` // Всего пройдено тестов

	User User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"` // Связь с пользователем
}
