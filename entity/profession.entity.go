package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Profession struct {
	Id            uuid.UUID        `json:"id" gorm:"column:id"`
	UserId        uuid.UUID        `json:"userId" gorm:"column:userId"`
	CategoryId    uuid.UUID        `json:"categoryId" gorm:"column:categoryId"`
	SubCategoryId uuid.UUID        `json:"subcategoryId" gorm:"column:subcategoryId"`
	Category      *Category        `json:"category" gorm:"foreignKey:categoryId"`
	Subcategory   *Subcategory     `json:"subcategory" gorm:"foreignKey:subcategoryId"`
	ProofMediaUrl string           `json:"proofMediaUrl" gorm:"column:proofMediaUrl"`
	Status        ProfessionStatus `json:"status" gorm:"column:status"`
	//Active        bool             `json:"active" gorm:"column:active;default:false"`
	IsShow bool `json:"-" gorm:"column:isShow;default:true"`
	BaseEntity
}

type ProfessionStatus string

const (
	PROFESSION_STATUS_IS_WAITING  ProfessionStatus = "IS_WAITING"
	PROFESSION_STATUS_IS_ACCEPTED ProfessionStatus = "IS_ACCEPTED"
	PROFESSION_STATUS_IS_DECLINED ProfessionStatus = "IS_DECLINED"
)

func (e *Profession) BeforeCreate(tx *gorm.DB) (err error) {
	e.Id = uuid.New()
	return
}
