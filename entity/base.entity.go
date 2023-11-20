package entity

import (
	"time"

	"gorm.io/gorm"
)

type BaseEntity struct {
	CreatedAt int64 `json:"createdAt" gorm:"column:createdAt;autoUpdateTime:milli"`
	UpdatedAt int64 `json:"updatedAt" gorm:"column:updatedAt;autoUpdateTime:milli"`
	DeletedAt int64 `json:"-" gorm:"column:deletedAt;index;default:null"`
}

type BaseEntityTime struct {
	CreatedAt time.Time      `json:"createdAt" gorm:"column:createdAt"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"column:updatedAt"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"column:deletedAt;index"`
}
