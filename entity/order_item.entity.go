package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Order Item is divided into it's own object if the events are not continuous or not the same product

type OrderItem struct {
	Id            uuid.UUID     `json:"id" gorm:"column:id"`
	OrderId       uuid.UUID     `json:"orderId" gorm:"column:orderId"`
	ProductId     uuid.UUID     `json:"productId" gorm:"column:productId"`
	UserId        uuid.UUID     `json:"userId" gorm:"column:userId"`
	Amount        int64         `json:"amount" gorm:"column:amount"`
	Starttime     int64         `json:"starttime" gorm:"column:starttime"`
	Endtime       int64         `json:"endtime" gorm:"column:endtime"`
	TotalPrice    int64         `json:"totalPrice" gorm:"column:totalPrice"`
	Product       *Product      `json:"product" gorm:"foreignKey:productId"`
	Status        OrderStatus   `json:"status" gorm:"column:status"`
	PaymentMethod PaymentMethod `json:"paymentMethod" gorm:"column:paymentMethod"`
	Version       int64         `json:"version" gorm:"column:version"`
	BaseEntity
}

type OrderStatus string

const (
	ORDER_STATUS_ON_WAITING_FOR_PAYMENT              OrderStatus = "ON_WAITING_FOR_PAYMENT"
	ORDER_STATUS_ON_WAITING_CONFIRMATION             OrderStatus = "ON_WAITING_CONFIRMATION"
	ORDER_STATUS_ON_SCHEDULE                         OrderStatus = "ON_SCHEDULE"
	ORDER_STATUS_ON_GOING                            OrderStatus = "ON_GOING"
	ORDER_STATUS_ON_REVIEW                           OrderStatus = "ON_REVIEW"
	ORDER_STATUS_ON_REPORTED_PROCESSED_BY_CONSULTANT OrderStatus = "ON_REPORTED_PROCESSED_BY_CONSULTANT"
	ORDER_STATUS_ON_REPORTED_PROCESSED_BY_ADMIN      OrderStatus = "ON_REPORTED_PROCESSED_BY_ADMIN"
	ORDER_STATUS_COMPLETED                           OrderStatus = "COMPLETED"
	ORDER_STATUS_CANCELED                            OrderStatus = "CANCELED"
	ORDER_STATUS_REFUNDED                            OrderStatus = "REFUNDED"
)

type PaymentMethod string

const (
	PAYMENT_METHOD_GOPAY PaymentMethod = "GOPAY"
)

func (e *OrderItem) BeforeCreate(tx *gorm.DB) (err error) {
	e.Id = uuid.New()
	return
}
