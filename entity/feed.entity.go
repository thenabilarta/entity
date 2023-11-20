package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Feed struct {
	Id            uuid.UUID      `json:"id" gorm:"column:id"`
	Title         string         `json:"title" gorm:"column:title"`
	Content       string         `json:"content" gorm:"column:content"`
	MediaUrl      string         `json:"mediaUrl" gorm:"column:mediaUrl"`
	Likes         int64          `json:"likes" gorm:"column:likes"`
	Views         int64          `json:"views" gorm:"column:views"`
	Rating        float64        `json:"rating" gorm:"column:rating"`
	TotalComment  int64          `json:"totalComments" gorm:"column:totalComments"`
	Tag           string         `json:"tag" gorm:"column:tag"`
	UserId        uuid.UUID      `json:"-" gorm:"column:userId"`
	CategoryId    uuid.UUID      `json:"-" gorm:"column:categoryId"`
	SubcategoryId uuid.UUID      `json:"-" gorm:"column:subcategoryId"`
	User          *User          `json:"user" gorm:"foreignKey:userId"`
	Category      *Category      `json:"category" gorm:"foreignKey:categoryId"`
	Subcategory   *Subcategory   `json:"subcategory" gorm:"foreignKey:subcategoryId"`
	Comment       *[]FeedComment `json:"comment" gorm:"foreignKey:feedId"`
	IsShow        bool           `json:"-" gorm:"column:isShow;default:true"`
	BaseEntity
}

func (e *Feed) BeforeCreate(tx *gorm.DB) (err error) {
	e.Id = uuid.New()

	return nil
}
