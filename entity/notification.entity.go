package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Notification struct {
	Id             uuid.UUID      `json:"id" gorm:"column:id"`
	UserId         uuid.NullUUID  `json:"-" gorm:"column:userId"`
	NotifationType NotifationType `json:"notificationType" gorm:"column:notificationType"`
	Title          string         `json:"title" gorm:"column:title"`
	Message        string         `json:"message" gorm:"column:message"`
	Seen           bool           `json:"seen" gorm:"column:seen"`
	BaseEntity
}

func (e *Notification) BeforeCreate(tx *gorm.DB) (err error) {
	e.Id = uuid.New()
	return
}

type NotifationType string

const (
	NOTIFICATION_TYPE_ORDER NotifationType = "NOTIFICATION_TYPE_ORDER"
)
