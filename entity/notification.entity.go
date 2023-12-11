package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Notification struct {
	Id             uuid.UUID      `json:"id" gorm:"column:id"`
	UserId         string         `json:"userId" gorm:"column:userId"`
	ToUserId       string         `json:"toUserId" gorm:"column:toUserId"`
	NotifationType NotifationType `json:"notificationType" gorm:"column:notificationType"`
	ItemId         string         `json:"itemId" gorm:"column:itemId"`
	ItemType       ItemType       `json:"itemType" gorm:"column:itemType"`
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
	NOTIFICATION_ORDER_WAITING_FOR_PAYMENT      NotifationType = "NOTIFICATION_ORDER_WAITING_FOR_PAYMENT"
	NOTIFICATION_ORDER_WAITING_FOR_CONFIRMATION NotifationType = "NOTIFICATION_ORDER_WAITING_FOR_CONFIRMATION"
	NOTIFICATION_ORDER_ON_SCHEDULE              NotifationType = "NOTIFICATION_ORDER_ON_SCHEDULE"
	NOTIFICATION_ORDER_ON_GOING                 NotifationType = "NOTIFICATION_ORDER_ON_GOING"
	NOTIFICATION_ORDER_FINISH                   NotifationType = "NOTIFICATION_ORDER_FINISH"
	NOTIFICATION_ORDER_GET_REPORT               NotifationType = "NOTIFICATION_ORDER_GET_REPORT"

	NOTIFICATION_ARTICLE_CREATED     NotifationType = "NOTIFICATION_ARTICLE_CREATED"
	NOTIFICATION_ARTICLE_GET_COMMENT NotifationType = "NOTIFICATION_ARTICLE_GET_COMMENT"
	NOTIFICATION_ARTICLE_GET_LIKE    NotifationType = "NOTIFICATION_ARTICLE_GET_LIKE"

	NOTIFICATION_QUESTION_CREATED            NotifationType = "NOTIFICATION_QUESTION_CREATED"
	NOTIFICATION_QUESTION_LIKED              NotifationType = "NOTIFICATION_QUESTION_LIKED"
	NOTIFICATION_QUESTION_GET_ANSWER         NotifationType = "NOTIFICATION_QUESTION_GET_ANSWER"
	NOTIFICATION_QUESTION_ANSWER_GET_COMMENT NotifationType = "NOTIFICATION_QUESTION_ANSWER_GET_COMMENT"

	NOTIFICATION_REVIEW_CREATED NotifationType = "NOTIFICATION_REVIEW_CREATED"

	NOTIFICATION_PROFESSION_REQUEST_CREATED  NotifationType = "NOTIFICATION_PROFESSION_REQUEST_CREATED"
	NOTIFICATION_PROFESSION_REQUEST_DECLINED NotifationType = "NOTIFICATION_PROFESSION_REQUEST_DECLINED"
	NOTIFICATION_PROFESSION_REQUEST_ACCEPTED NotifationType = "NOTIFICATION_PROFESSION_REQUEST_ACCEPTED"

	NOTIFICATION_SERVICES_REQUEST_CREATED  NotifationType = "NOTIFICATION_SERVICES_REQUEST_CREATED"
	NOTIFICATION_SERVICES_REQUEST_DECLINED NotifationType = "NOTIFICATION_SERVICES_REQUEST_DECLINED"
	NOTIFICATION_SERVICES_REQUEST_ACCEPTED NotifationType = "NOTIFICATION_SERVICES_REQUEST_ACCEPTED"

	NOTIFICATION_SERVICES_NEW_CONSULTANT_REQUEST  NotifationType = "NOTIFICATION_SERVICES_NEW_CONSULTANT_REQUEST"
	NOTIFICATION_SERVICES_NEW_CONSULTANT_DECLINED NotifationType = "NOTIFICATION_SERVICES_NEW_CONSULTANT_DECLINED"
	NOTIFICATION_SERVICES_NEW_CONSULTANT_ACCEPTED NotifationType = "NOTIFICATION_SERVICES_NEW_CONSULTANT_ACCEPTED"

	NOTIFICATION_SERVICES_CONSULTANT_CHANGE_PROFILE_REQUEST  NotifationType = "NOTIFICATION_SERVICES_CONSULTANT_CHANGE_PROFILE_REQUEST"
	NOTIFICATION_SERVICES_CONSULTANT_CHANGE_PROFILE_DECLINED NotifationType = "NOTIFICATION_SERVICES_CONSULTANT_CHANGE_PROFILE_DECLINED"
	NOTIFICATION_SERVICES_CONSULTANT_CHANGE_PROFILE_ACCEPTED NotifationType = "NOTIFICATION_SERVICES_CONSULTANT_CHANGE_PROFILE_ACCEPTED"
)

type ItemType string

const (
	NOTIFICATION_ITEM_TYPE_ARTICLE                           ItemType = "NOTIFICATION_ITEM_TYPE_ARTICLE"
	NOTIFICATION_ITEM_TYPE_ASKME                             ItemType = "NOTIFICATION_ITEM_TYPE_ASKME"
	NOTIFICATION_ITEM_TYPE_ORDER                             ItemType = "NOTIFICATION_ITEM_TYPE_ORDER"
	NOTIFICATION_ITEM_TYPE_SERVICES_REQUEST                  ItemType = "NOTIFICATION_ITEM_TYPE_SERVICES_REQUEST"
	NOTIFICATION_ITEM_TYPE_PROFESSION_REQUEST                ItemType = "NOTIFICATION_ITEM_TYPE_PROFESSION_REQUEST"
	NOTIFICATION_ITEM_TYPE_NEW_CONSULTANT_REQUEST            ItemType = "NOTIFICATION_ITEM_TYPE_NEW_CONSULTANT_REQUEST"
	NOTIFICATION_ITEM_TYPE_CONSULTANT_CHANGE_PROFILE_REQUEST ItemType = "NOTIFICATION_ITEM_TYPE_CONSULTANT_CHANGE_PROFILE_REQUEST"
)
