package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	Id            uuid.UUID   `json:"id" gorm:"column:id"`
	Title         string      `json:"title" gorm:"column:title"`
	Description   string      `json:"description" gorm:"column:description"`
	ImageUrl      string      `json:"imageUrl" gorm:"column:imageUrl"`
	Price         int64       `json:"price" gorm:"column:price"`
	Duration      int64       `json:"duration" gorm:"column:duration"`
	TotalSold     int64       `json:"totalSold" gorm:"column:totalSold"`
	Likes         int64       `json:"likes" gorm:"column:likes"`
	Views         int64       `json:"views" gorm:"column:views"`
	Rating        float64     `json:"rating" gorm:"column:rating"`
	Tag           string      `json:"tag" gorm:"column:tag"`
	Schedule      *[]Schedule `json:"schedule" gorm:"foreignKey:productId"`
	UserId        uuid.UUID   `json:"userId" gorm:"column:userId"`
	CategoryId    uuid.UUID   `json:"categoryId" gorm:"column:categoryId"`
	SubcategoryId uuid.UUID   `json:"subcategoryId" gorm:"column:subcategoryId"`
	ProfessionId  uuid.UUID   `json:"professionId" gorm:"column:professionId"`
	IsShow        bool        `json:"isShow" gorm:"column:isShow;default:true"`
	BaseEntity
}

func (e *Product) BeforeCreate(tx *gorm.DB) (err error) {
	e.Id = uuid.New()
	return
}

type Schedule struct {
	Id        uuid.UUID    `json:"id" gorm:"column:id"`
	ProductId uuid.UUID    `json:"productId" gorm:"column:productId"`
	Day       Day          `json:"day" gorm:"column:day"`
	StartTime ScheduleTime `json:"starttime" gorm:"column:starttime"`
	EndTime   ScheduleTime `json:"endtime" gorm:"column:endtime"`
	BaseEntity
}

func (e *Schedule) BeforeCreate(tx *gorm.DB) (err error) {
	e.Id = uuid.New()
	return
}

type Day string

const (
	SCHEDULE_DAY_MONDAY    Day = "MONDAY"
	SCHEDULE_DAY_TUESDAY   Day = "TUESDAY"
	SCHEDULE_DAY_WEDNESDAY Day = "WEDNESDAY"
	SCHEDULE_DAY_THURSDAY  Day = "THURSDAY"
	SCHEDULE_DAY_FRIDAY    Day = "FRIDAY"
	SCHEDULE_DAY_SATURDAY  Day = "SATURDAY"
	SCHEDULE_DAY_SUNDAY    Day = "SUNDAY"
)

type ScheduleTime string

const (
	SCHEDULETIME_0000 ScheduleTime = "0000"
	SCHEDULETIME_0015 ScheduleTime = "0015"
	SCHEDULETIME_0030 ScheduleTime = "0030"
	SCHEDULETIME_0045 ScheduleTime = "0045"

	SCHEDULETIME_0100 ScheduleTime = "0100"
	SCHEDULETIME_0115 ScheduleTime = "0115"
	SCHEDULETIME_0130 ScheduleTime = "0130"
	SCHEDULETIME_0145 ScheduleTime = "0145"

	SCHEDULETIME_0200 ScheduleTime = "0200"
	SCHEDULETIME_0215 ScheduleTime = "0215"
	SCHEDULETIME_0230 ScheduleTime = "0230"
	SCHEDULETIME_0245 ScheduleTime = "0245"

	SCHEDULETIME_0300 ScheduleTime = "0300"
	SCHEDULETIME_0315 ScheduleTime = "0315"
	SCHEDULETIME_0330 ScheduleTime = "0330"
	SCHEDULETIME_0345 ScheduleTime = "0345"

	SCHEDULETIME_0400 ScheduleTime = "0400"
	SCHEDULETIME_0415 ScheduleTime = "0415"
	SCHEDULETIME_0430 ScheduleTime = "0430"
	SCHEDULETIME_0445 ScheduleTime = "0445"

	SCHEDULETIME_0500 ScheduleTime = "0500"
	SCHEDULETIME_0515 ScheduleTime = "0515"
	SCHEDULETIME_0530 ScheduleTime = "0530"
	SCHEDULETIME_0545 ScheduleTime = "0545"

	SCHEDULETIME_0600 ScheduleTime = "0600"
	SCHEDULETIME_0615 ScheduleTime = "0615"
	SCHEDULETIME_0630 ScheduleTime = "0630"
	SCHEDULETIME_0645 ScheduleTime = "0645"

	SCHEDULETIME_0700 ScheduleTime = "0700"
	SCHEDULETIME_0715 ScheduleTime = "0715"
	SCHEDULETIME_0730 ScheduleTime = "0730"
	SCHEDULETIME_0745 ScheduleTime = "0745"

	SCHEDULETIME_0800 ScheduleTime = "0800"
	SCHEDULETIME_0815 ScheduleTime = "0815"
	SCHEDULETIME_0830 ScheduleTime = "0830"
	SCHEDULETIME_0845 ScheduleTime = "0845"

	SCHEDULETIME_0900 ScheduleTime = "0900"
	SCHEDULETIME_0915 ScheduleTime = "0915"
	SCHEDULETIME_0930 ScheduleTime = "0930"
	SCHEDULETIME_0945 ScheduleTime = "0945"

	SCHEDULETIME_1000 ScheduleTime = "1000"
	SCHEDULETIME_1015 ScheduleTime = "1015"
	SCHEDULETIME_1030 ScheduleTime = "1030"
	SCHEDULETIME_1045 ScheduleTime = "1045"

	SCHEDULETIME_1100 ScheduleTime = "1100"
	SCHEDULETIME_1115 ScheduleTime = "1115"
	SCHEDULETIME_1130 ScheduleTime = "1130"
	SCHEDULETIME_1145 ScheduleTime = "1145"

	SCHEDULETIME_1200 ScheduleTime = "1200"
	SCHEDULETIME_1215 ScheduleTime = "1215"
	SCHEDULETIME_1230 ScheduleTime = "1230"
	SCHEDULETIME_1245 ScheduleTime = "1245"

	SCHEDULETIME_1300 ScheduleTime = "1300"
	SCHEDULETIME_1315 ScheduleTime = "1315"
	SCHEDULETIME_1330 ScheduleTime = "1330"
	SCHEDULETIME_1345 ScheduleTime = "1345"

	SCHEDULETIME_1400 ScheduleTime = "1400"
	SCHEDULETIME_1415 ScheduleTime = "1415"
	SCHEDULETIME_1430 ScheduleTime = "1430"
	SCHEDULETIME_1445 ScheduleTime = "1445"

	SCHEDULETIME_1500 ScheduleTime = "1500"
	SCHEDULETIME_1515 ScheduleTime = "1515"
	SCHEDULETIME_1530 ScheduleTime = "1530"
	SCHEDULETIME_1545 ScheduleTime = "1545"

	SCHEDULETIME_1600 ScheduleTime = "1600"
	SCHEDULETIME_1615 ScheduleTime = "1615"
	SCHEDULETIME_1630 ScheduleTime = "1630"
	SCHEDULETIME_1645 ScheduleTime = "1645"

	SCHEDULETIME_1700 ScheduleTime = "1700"
	SCHEDULETIME_1715 ScheduleTime = "1715"
	SCHEDULETIME_1730 ScheduleTime = "1730"
	SCHEDULETIME_1745 ScheduleTime = "1745"

	SCHEDULETIME_1800 ScheduleTime = "1800"
	SCHEDULETIME_1815 ScheduleTime = "1815"
	SCHEDULETIME_1830 ScheduleTime = "1830"
	SCHEDULETIME_1845 ScheduleTime = "1845"

	SCHEDULETIME_1900 ScheduleTime = "1900"
	SCHEDULETIME_1915 ScheduleTime = "1915"
	SCHEDULETIME_1930 ScheduleTime = "1930"
	SCHEDULETIME_1945 ScheduleTime = "1945"

	SCHEDULETIME_2000 ScheduleTime = "2000"
	SCHEDULETIME_2015 ScheduleTime = "2015"
	SCHEDULETIME_2030 ScheduleTime = "2030"
	SCHEDULETIME_2045 ScheduleTime = "2045"

	SCHEDULETIME_2100 ScheduleTime = "2100"
	SCHEDULETIME_2115 ScheduleTime = "2115"
	SCHEDULETIME_2130 ScheduleTime = "2130"
	SCHEDULETIME_2145 ScheduleTime = "2145"

	SCHEDULETIME_2200 ScheduleTime = "2200"
	SCHEDULETIME_2215 ScheduleTime = "2215"
	SCHEDULETIME_2230 ScheduleTime = "2230"
	SCHEDULETIME_2245 ScheduleTime = "2245"

	SCHEDULETIME_2300 ScheduleTime = "2300"
	SCHEDULETIME_2315 ScheduleTime = "2315"
	SCHEDULETIME_2330 ScheduleTime = "2330"
	SCHEDULETIME_2345 ScheduleTime = "2345"
)
