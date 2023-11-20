package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Order struct {
	Id            uuid.UUID     `json:"id" gorm:"column:id"`
	UserId        uuid.NullUUID `json:"-" gorm:"column:userId"`
	TotalPrice    int64         `json:"totalPrice" gorm:"column:totalPrice"`
	IsPaid        bool          `json:"isPaid" gorm:"column:isPaid"`
	PaymentId     string        `json:"paymentId" gorm:"column:paymentId"`
	PaymentData   string        `json:"paymentData" gorm:"column:paymentData"`
	User          *User         `json:"user" gorm:"foreignKey:userId"`
	Version       int64         `json:"version" gorm:"column:version"`
	OrderItem     *[]OrderItem  `json:"order" gorm:"foreignKey:orderId"`
	Status        OrderStatus   `json:"status" gorm:"column:status"`
	PaymentMethod PaymentMethod `json:"paymentMethod" gorm:"column:paymentMethod"`
	BaseEntity
}

func (e *Order) BeforeCreate(tx *gorm.DB) (err error) {
	e.Id = uuid.New()
	return
}
