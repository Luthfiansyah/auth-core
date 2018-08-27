package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Client struct {
	gorm.Model
	ID           uint      `gorm:"-; primary_key"`
	Name         string    `gorm:"column:name; index:client_name; not null" json:"name"`
	Username     string    `gorm:"column:username; index:client_user_username; not null" json:"username"`
	Password     string    `gorm:"column:password; index:client_user_password; not null" json:"password"`
	ClientTypeID uint      `gorm:"column:client_type_id; index:client_token_client_type_id; not null" json:"client_type_id"`
	Description  string    `gorm:"column:description; index:client_description; not null" json:"description"`
	CreatedBy    uint      `gorm:"column:created_by; index:client_user_created_by" json:"created_by"`
	UpdatedBy    uint      `gorm:"column:updated_by; index:client_user_updated_by" json:"updated_by"`
	CreatedAt    time.Time `gorm:"-" json:"created_at"`
	UpdatedAt    time.Time `gorm:"-" json:"update_at"`
	DeletedAt    time.Time `gorm:"-" json:"deleted_at"`
	RowStatus    int32     `gorm:"column:row_status; index:client_row_status; default:1" json:"row_status"`
}

type ClientAuth struct {
	ID             NullInt64  `db:"id" json:"id"`
	Name           NullString `db:"name" json:"name"`
	Username       NullString `db:"username" json:"username"`
	Password       NullString `db:"password" json:"password"`
	Description    NullString `db:"description" json:"description"`
	ClientTypeID   NullInt64  `db:"client_type_id" json:"client_type_id"`
	ClientTypeName NullString `db:"client_type_name" json:"client_type_name"`
	CreatedBy      NullInt64  `db:"created_by" json:"created_by"`
	UpdatedBy      NullInt64  `db:"updated_by" json:"updated_by"`
	CreatedAt      NullTime   `db:"created_at" json:"created_at"`
	UpdatedAt      NullTime   `db:"updated_at" json:"updated_at"`
	DeletedAt      NullTime   `db:"delete_at" json:"delete_at"`
	RowStatus      NullInt64  `db:"row_status" json:"row_status"`
}
