package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FeedComment struct {
	Id      uuid.UUID     `json:"id" gorm:"column:id"`
	Content string        `json:"content" gorm:"column:content"`
	UserId  uuid.NullUUID `json:"-" gorm:"column:userId"`
	FeedId  uuid.UUID     `json:"feedId" gorm:"column:feedId"`
	User    *User         `json:"user" gorm:"foreignKey:userId"`
	IsShow  bool          `json:"-" gorm:"column:isShow;default:true"`
	BaseEntity
}

func (e *FeedComment) BeforeCreate(tx *gorm.DB) (err error) {
	e.Id = uuid.New()
	return
}
