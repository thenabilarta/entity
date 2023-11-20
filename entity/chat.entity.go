package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Chat struct {
	Id            uuid.UUID     `json:"id" gorm:"column:id"`
	UserId_0      uuid.UUID     `json:"userId_0" gorm:"column:userId_0"`
	UserId_1      uuid.UUID     `json:"userId_1" gorm:"column:userId_1"`
	IsAppointment bool          `json:"isAppointment" gorm:"column:isAppointment"`
	OrderItemId   uuid.NullUUID `json:"orderItemId" gorm:"column:orderItemId"`
	User_0        *User         `json:"user_0" gorm:"foreignKey:userId_0"`
	User_1        *User         `json:"user_1" gorm:"foreignKey:userId_1"`
	OrderItem     *OrderItem    `json:"orderItem" gorm:"foreignKey:orderItemId"`
	Version       int64         `json:"version" gorm:"column:version"`
	ChatItem      *[]ChatItem   `json:"chatItem" gorm:"foreignKey:chatId"`
	BaseEntity
}

func (e *Chat) BeforeCreate(tx *gorm.DB) (err error) {
	e.Id = uuid.New()
	return
}
