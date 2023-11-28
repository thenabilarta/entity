package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Subcategory struct {
	Id         uuid.UUID     `json:"id" gorm:"column:id"`
	Name       string        `json:"name" gorm:"column:name"`
	CategoryId uuid.NullUUID `json:"categoryId" gorm:"column:categoryId"`
	PriceRange string        `json:"priceRange" gorm:"column:priceRange"`
	ImageUrl   string        `json:"imageUrl" gorm:"column:imageUrl"`
	IsShow     bool          `json:"-" gorm:"column:isShow;default:true"`
	BaseEntity
}

func (e *Subcategory) BeforeCreate(tx *gorm.DB) (err error) {
	e.Id = uuid.New()
	return
}
