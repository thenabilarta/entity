package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ArticleComment struct {
	Id        uuid.UUID `json:"id" gorm:"column:id"`
	Content   string    `json:"content" gorm:"column:content"`
	UserId    uuid.UUID `json:"userId" gorm:"column:userId"`
	ArticleId uuid.UUID `json:"articleId" gorm:"column:articleId"`
	IsShow    bool      `json:"-" gorm:"column:isShow;default:true"`
	BaseEntity
}

func (e *ArticleComment) BeforeCreate(tx *gorm.DB) (err error) {
	e.Id = uuid.New()
	return
}
