package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	Id             uuid.UUID `json:"id" gorm:"column:id"`
	Name           string    `json:"name" gorm:"column:name"`
	Email          string    `json:"email" gorm:"column:email"`
	Phone          string    `json:"phone" gorm:"column:phone"`
	HashedPassword string    `json:"-" gorm:"column:hashedPassword"`
	Balance        int64     `json:"balance" gorm:"column:balance"`
	Role           UserRole  `json:"role" gorm:"column:role"`
	ProfileDesc    string    `json:"profileDesc" gorm:"column:profileDesc"`
	ProfilePic     string    `json:"profilePic" gorm:"column:profilePic"`
	OnlineStatus   bool      `json:"onlineStatus" gorm:"column:onlineStatus"`
	Status         bool      `json:"status" gorm:"column:status"`
	EmailVerified  int64     `json:"emailVerified" gorm:"column:emailVerified"`
	PhoneVerified  int64     `json:"phoneVerified" gorm:"column:EmailVerified"`
	Active         bool      `json:"active" gorm:"column:active"`
	BaseEntity
}

type UserConsultantProfile struct {
	Id         uuid.UUID `json:"id" gorm:"column:id"`
	Name       string    `json:"name" gorm:"column:name"`
	ProfilePic string    `json:"profilePic" gorm:"column:profilePic"`
	Role       UserRole  `json:"role" gorm:"column:role"`
}

func (e *User) BeforeCreate(tx *gorm.DB) (err error) {
	e.Id = uuid.New()
	return
}

type UserRole string

const (
	USER_ROLE_CUSTOMER   UserRole = "USER_ROLE_CUSTOMER"
	USER_ROLE_CONSULTANT UserRole = "USER_ROLE_CONSULTANT"
	USER_ROLE_ADMIN      UserRole = "USER_ROLE_ADMIN"
)
