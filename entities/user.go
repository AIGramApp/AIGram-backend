package entities

import "time"

// User model
type User struct {
	ID       int64   `gorm:"id;primary_key;auto_increment:false;not null" json:"id"`
	Email    string  `gorm:"email;unique" json:"email"`
	Name     string  `gorm:"name;not null" json:"name"`
	Avatar   *string `gorm:"avatar" json:"avatar"`
	Username string  `gorm:"username;not null" json:"username"`

	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `sql:"index" json:"-"`
}
