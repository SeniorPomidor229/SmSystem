package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID         uuid.UUID `gorm:"type:uuid"`
	Username   string    `json:"username,omitempty" gorm:"column:username"`
	Password   string    `json:"password,omitempty" gorm:"column:password"`
	TelegramId string    `json:"telegram_id,omitempty" gorm:"column:telegram_id"`
}
