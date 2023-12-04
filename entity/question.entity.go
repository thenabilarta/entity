package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Question struct {
	Id          uuid.UUID `json:"id" gorm:"column:id"`
	Content     string    `json:"content" gorm:"column:content"`
	Likes       int64     `json:"likes" gorm:"column:likes"`
	TotalAnswer int64     `json:"totalAnswer" gorm:"column:totalAnswer"`
	Tag         string    `json:"tag" gorm:"column:tag"`
	UserId      uuid.UUID `json:"userId" gorm:"column:userId"`
	User        *User     `json:"user" gorm:"foreignKey:userId"`
	Answer      *[]Answer `json:"answer" gorm:"foreignKey:questionId"`
	IsShow      bool      `json:"-" gorm:"column:isShow;default:true"`
	BaseEntity
}

func (e *Question) BeforeCreate(tx *gorm.DB) (err error) {
	e.Id = uuid.New()
	return
}
