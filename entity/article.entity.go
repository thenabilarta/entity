package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Article struct {
	Id            uuid.UUID         `json:"id" gorm:"column:id"`
	Title         string            `json:"title" gorm:"column:title"`
	Content       string            `json:"content" gorm:"column:content"`
	Likes         int64             `json:"likes" gorm:"column:likes"`
	Views         int64             `json:"views" gorm:"column:views"`
	Rating        float64           `json:"rating" gorm:"column:rating"`
	TotalComments int64             `json:"totalComments" gorm:"column:totalComments"`
	Tag           string            `json:"tag" gorm:"column:tag"`
	UserId        uuid.UUID         `json:"userId" gorm:"column:userId"`
	CategoryId    uuid.UUID         `json:"categoryId" gorm:"column:categoryId"`
	SubcategoryId uuid.UUID         `json:"subcategoryId" gorm:"column:subcategoryId"`
	Comment       *[]ArticleComment `json:"comment" gorm:"foreignKey:articleId"`
	Category      Category          `json:"category" gorm:"foreignKey:categoryId"`
	Subcategory   Subcategory       `json:"subcategory" gorm:"foreignKey:subcategoryId"`
	MediaUrl      string            `json:"mediaUrl" gorm:"column:mediaUrl"`
	IsShow        bool              `json:"-" gorm:"column:isShow;default:true"`
	BaseEntity
}

func (e *Article) BeforeCreate(tx *gorm.DB) (err error) {
	e.Id = uuid.New()
	return
}
