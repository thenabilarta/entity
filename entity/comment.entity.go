package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Comment struct {
	Id       uuid.UUID     `json:"id" gorm:"column:id"`
	Content  string        `json:"content" gorm:"column:content"`
	UserId   uuid.NullUUID `json:"-" gorm:"column:userId"`
	AnswerId uuid.UUID     `json:"answerId" gorm:"column:answerId"`
	User     *User         `json:"user" gorm:"foreignKey:userId"`
	IsShow   bool          `json:"-" gorm:"column:isShow;default:true"`
	BaseEntity
}

func (e *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	e.Id = uuid.New()
	return
}
