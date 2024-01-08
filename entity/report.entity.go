package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Report struct {
	Id           uuid.UUID        `json:"id" gorm:"column:id"`
	ConsultantId uuid.UUID        `json:"consultantId" gorm:"column:consultantId"`
	CustomerId   uuid.UUID        `json:"customerId" gorm:"column:customerId"`
	OrderItemId  uuid.UUID        `json:"orderItemId" gorm:"column:orderItemId"`
	Value        int              `json:"value" gorm:"column:value"`
	Title        string           `json:"title" gorm:"column:title"`
	Description  string           `json:"description" gorm:"column:description"`
	ImageUrl     string           `json:"imageUrl" gorm:"column:imageUrl"`
	Status       ReportStatus     `json:"status" gorm:"column:status"`
	Resolution   ReportResolution `json:"resolution" gorm:"column:resolution"`
	BaseEntity
}

func (e *Report) BeforeCreate(tx *gorm.DB) (err error) {
	e.Id = uuid.New()
	return
}

type ReportStatus string

const (
	REPORT_STATUS_CREATED     ReportStatus = "REPORT_STATUS_CREATED"
	REPORT_STATUS_ON_PROGRESS ReportStatus = "REPORT_STATUS_ON_PROGRESS"
	REPORT_STATUS_FINISHED    ReportStatus = "REPORT_STATUS_FINISHED"
)

type ReportResolution string

const (
	REPORT_RESOLUTION_APPROVED    ReportResolution = "REPORT_RESOLUTION_APPROVED"
	REPORT_RESOLUTION_DECLINED    ReportResolution = "REPORT_RESOLUTION_DECLINED"
	REPORT_RESOLUTION_ON_PROGRESS ReportResolution = "REPORT_RESOLUTION_ON_PROGRESS"
)
