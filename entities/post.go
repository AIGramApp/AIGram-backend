package entities

import (
	"time"

	"github.com/lib/pq"
)

// Post published by users
type Post struct {
	ID          int64          `gorm:"id;primary_key;not null" json:"id"`
	Title       string         `json:"title" binding:"required"`
	Description string         `json:"description" binding:"required"`
	Image       string         `json:"image" binding:"required"`
	Link        string         `json:"link" binding:"required"`
	Tags        pq.StringArray `gorm:"type:varchar(100)[]" json:"tags" binding:"required"`
	User        User           `gorm:"foreignkey:UserRefer" json:"user"`
	UserRefer   int64          `json:"-"`
	CreatedAt   time.Time      `json:"date"`
	UpdatedAt   time.Time      `json:"-"`
	DeletedAt   *time.Time     `sql:"index" json:"-"`
}
