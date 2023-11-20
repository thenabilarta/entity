package entity

import (
	"github.com/google/uuid"
)

type OfficerRefreshToken struct {
	Id        uuid.UUID `gorm:"column:id"`
	UserId    uuid.UUID `gorm:"column:userId"`
	Hash      string    `gorm:"column:hash"`
	ExpiredAt int64     `gorm:"column:expiredAt"`
	BaseEntity
}
