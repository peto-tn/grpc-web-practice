package entity

import (
	"time"
)

type Article struct {
	Id        int64
	Title     string
	Body      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
