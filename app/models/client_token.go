package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type ClientToken struct {
	gorm.Model
	ID        uint      `gorm:"-; primary_key"`
	ClientID  uint      `gorm:"column:client_id; index:client_token_client_id; not null" json:"client_id"`
	Token     string    `gorm:"column:token; index:token; not null" json:"token"`
	ExpiredAt int32     `gorm:"column:expired_at; index:client_token_expired_at; not null" json:"expired_at"`
	CreatedBy uint      `gorm:"column:created_by; index:client_token_user_created_by" json:"created_by"`
	UpdatedBy uint      `gorm:"column:updated_by; index:client_token_user_updated_by" json:"updated_by"`
	CreatedAt time.Time `gorm:"-" json:"created_at"`
	UpdatedAt time.Time `gorm:"-" json:"update_at"`
	DeletedAt time.Time `gorm:"-" json:"deleted_at"`
	RowStatus int32     `gorm:"column:row_status; index:client_token_row_status; default:1" json:"row_status"`
}

type ClientAuthToken struct {
	ID        NullInt64  `db:"id" json:"id"`
	Token     NullString `db:"token" json:"token"`
	CreatedBy NullInt64  `db:"created_by" json:"created_by"`
	UpdatedBy NullInt64  `db:"updated_by" json:"updated_by"`
	CreatedAt NullTime   `db:"created_at" json:"created_at"`
	UpdatedAt NullTime   `db:"updated_at" json:"updated_at"`
	DeletedAt NullTime   `db:"deleted_at" json:"deleted_at"`
	RowStatus NullInt64  `db:"row_status" json:"row_status"`
	ClientID  NullInt64  `db:"client_id" json:"client_id"`
}
