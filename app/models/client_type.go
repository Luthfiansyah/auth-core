package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type ClientType struct {
	gorm.Model
	ID          uint      `gorm:"-; primary_key"`
	Name        string    `gorm:"column:name; index:client_type_name; not null" json:"name"`
	Description string    `gorm:"column:description; index:client_type_description; not null" json:"description"`
	CreatedBy   uint      `gorm:"column:created_by; index:client_type_user_created_by" json:"created_by"`
	UpdatedBy   uint      `gorm:"column:updated_by; index:client_type_user_updated_by" json:"updated_by"`
	CreatedAt   time.Time `gorm:"-" json:"created_at"`
	UpdatedAt   time.Time `gorm:"-" json:"update_at"`
	DeletedAt   time.Time `gorm:"-" json:"deleted_at"`
	RowStatus   int32     `gorm:"column:row_status; index:client_type_row_status; default:1" json:"row_status"`
}
