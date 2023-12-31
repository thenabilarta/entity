package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Answer struct {
	Id           uuid.UUID              `json:"id" gorm:"column:id"`
	Content      string                 `json:"content" gorm:"column:content"`
	Likes        int64                  `json:"likes" gorm:"column:likes"`
	TotalComment int64                  `json:"totalComments" gorm:"column:totalComments"`
	Tag          string                 `json:"tag" gorm:"column:tag"`
	UserId       uuid.UUID              `json:"userId" gorm:"column:userId"`
	User         *UserConsultantProfile `json:"user" gorm:"foreignKey:userId"`
	QuestionId   uuid.UUID              `json:"questionId" gorm:"column:questionId"`
	Question     *Question              `json:"question" gorm:"foreignKey:questionId"`
	Comment      *[]Comment             `json:"comment" gorm:"foreignKey:answerId"`
	IsShow       bool                   `json:"isShow" gorm:"column:isShow;default:true"`
	BaseEntity
}

func (e *Answer) BeforeCreate(tx *gorm.DB) (err error) {
	e.Id = uuid.New()
	return
}
