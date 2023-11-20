package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ChatItem struct {
	Id        uuid.UUID `json:"id" gorm:"column:id"`
	ChatId    uuid.UUID `json:"chatId" gorm:"column:chatId"`
	Time      int64     `json:"time" gorm:"column:time"`
	Content   string    `json:"content" gorm:"column:content"`
	Direction bool      `json:"direction" gorm:"column:direction"`
	MediaType MediaType `json:"mediaType" gorm:"column:mediaType"`
	MediaUrl  string    `json:"mediaUrl" gorm:"column:mediaUrl"`
	Version   int64     `json:"version" gorm:"column:version"`
	BaseEntity
}

type MediaType string

const (
	CHAT_MEDIA_TYPE_IMAGE     MediaType = "MEDIA_TYPE_IMAGE"
	CHAT_MEDIA_TYPE_VIDEO     MediaType = "MEDIA_TYPE_VIDEO"
	CHAT_MEDIA_TYPE_TEXT_ONLY MediaType = "MEDIA_TYPE_TEXT_ONLY"
)

func (e *ChatItem) BeforeCreate(tx *gorm.DB) (err error) {
	e.Id = uuid.New()
	return
}
