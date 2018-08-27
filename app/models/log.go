package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Log struct {
	gorm.Model
	ID              uint      `gorm:"-"`
	Endpoint        string    `gorm:"column:endpoint; index:log_endpoint; not null" json:"endpoint"`
	RequestMessage  string    `gorm:"column:request_message;"`
	ResponseMessage string    `gorm:"column:response_message;"`
	RequestTime     time.Time `gorm:"column:request_time; index:log__request_time;"`
	ResponseTime    time.Time `gorm:"column:response_time; index:log__response_time;"`
	ElapsedTime     int32     `gorm:"column:elapsed_time; index:log__elapsed_time;"`
	CreatedBy       uint      `gorm:"column:created_by; index:log__created_by"`
	UpdatedBy       uint      `gorm:"column:updated_by; index:log__updated_by"`
	CreatedAt       time.Time `gorm:"-"`
	UpdatedAt       time.Time `gorm:"-"`
	DeletedAt       time.Time `gorm:"-"`
	RowStatus       int32     `gorm:"column:row_status; index:log__row_status;default:1"`
}
