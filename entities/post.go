package entities

import (
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

// Post published by users
type Post struct {
	gorm.Model
	Title       string         `json:"title" binding:"required"`
	Description string         `json:"description" binding:"required"`
	Image       string         `json:"image" binding:"required"`
	Link        string         `json:"link" binding:"required"`
	Tags        pq.StringArray `gorm:"type:varchar(100)[]" json:"tags" binding:"required"`
	User        User           `gorm:"foreignkey:UserRefer"`
	UserRefer   int64
}
