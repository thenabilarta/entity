package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Document struct {
	Id        uuid.UUID `json:"id" gorm:"column:id"`
	Name      string    `json:"name" gorm:"column:name"`
	UserId    string    `json:"userId" gorm:"column:userId"`
	MediaType string    `json:"mediaType" gorm:"column:mediaType"`
	MediaUrl  string    `json:"mediaUrl" gorm:"column:mediaUrl"`

	BaseEntity
}

func (e *Document) BeforeCreate(tx *gorm.DB) (err error) {
	e.Id = uuid.New()
	return
}
