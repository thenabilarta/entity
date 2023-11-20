package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Cart struct {
	Id        uuid.UUID `json:"id" gorm:"column:id"`
	ProductId uuid.UUID `json:"productId" gorm:"column:productId"`
	UserId    uuid.UUID `json:"userId" gorm:"column:userId"`
	Amount    int64     `json:"amount" gorm:"column:amount"`
	Starttime int64     `json:"starttime" gorm:"column:starttime"`
	Endtime   int64     `json:"endtime" gorm:"column:endtime"`
	User      *User     `json:"user" gorm:"foreignKey:userId"`
	Product   *Product  `json:"product" gorm:"foreignKey:productId"`
	Version   int64     `json:"version" gorm:"column:version"`
	BaseEntity
}

func (e *Cart) BeforeCreate(tx *gorm.DB) (err error) {
	e.Id = uuid.New()
	return
}
