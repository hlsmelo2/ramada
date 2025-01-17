package models

import (
	"time"
)

type Product struct {
	ID          uint64 `json:-`
	Name        string
	Description string
	Price       string
	Category    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
