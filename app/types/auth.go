package types

import "time"

type RequestGetToken struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ResponseGetToken struct {
	Token     string `json:"token"`
	ExpiredAt int64  `json:"expired_at"`
}

type InsertClientToken struct {
	ClientID  int64     `json:"client_id"`
	Token     string    `json:"token"`
	ExpiredAt int64     `json:"expired_at"`
	CreatedBy int64     `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
}

type UpdateClientToken struct {
	ClientID  int64     `json:"client_id"`
	DeletedAt time.Time `json:"expired_at"`
	UpdatedBy int64     `json:"updated_by"`
	UpdatedAt time.Time `json:"updated_at"`
}
