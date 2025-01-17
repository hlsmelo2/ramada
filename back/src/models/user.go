package models

import (
	"time"
)

type User struct {
	ID        uint64 `json:-`
	Name      string
	Username  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
