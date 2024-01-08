package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Feedback struct {
	Id           uuid.UUID `json:"id" gorm:"column:id"`
	ConsultantId uuid.UUID `json:"consultantId" gorm:"column:consultantId"`
	CustomerId   uuid.UUID `json:"customerId" gorm:"column:customerId"`
	OrderItemId  uuid.UUID `json:"orderItemId" gorm:"column:orderItemId"`
	Value        int       `json:"value" gorm:"column:value"`
	Description  string    `json:"description" gorm:"column:description"`
	BaseEntity
}

func (e *Feedback) BeforeCreate(tx *gorm.DB) (err error) {
	e.Id = uuid.New()
	return
}
