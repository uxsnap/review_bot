package entity

import (
	"gorm.io/gorm"
)

type ProblematicQuestion struct {
	gorm.Model
	WrongAnswersCount int `gorm:"default:0"` // Количество неправильных попыток

	User     User     `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`     // Связь с пользователем
	Question Question `gorm:"foreignKey:QuestionID;constraint:OnDelete:CASCADE"` // Связь с вопросом
}
