package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Category struct {
	Id          uuid.UUID     `json:"id" gorm:"column:id"`
	Name        string        `json:"name" gorm:"column:name"`
	Subcategory []Subcategory `json:"subcategory" gorm:"column:subcategory"`
	ImageUrl    string        `json:"imageUrl" gorm:"column:imageUrl"`
	IsShow      bool          `json:"-" gorm:"column:isShow;default:true"`
	BaseEntity
}

func (e *Category) BeforeCreate(tx *gorm.DB) (err error) {
	e.Id = uuid.New()
	return
}
