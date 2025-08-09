package models

import "time"

type User struct {
	ID        uint       `gorm:"primaryKey" json:"id" example:"1"`
	CreatedAt time.Time  `json:"created_at" example:"2025-08-09T18:15:06.000Z"`
	UpdatedAt time.Time  `json:"updated_at" example:"2025-08-09T18:15:06.000Z"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at,omitempty"`
	Username  string     `gorm:"unique" json:"username" example:"john_doe"`
	Password  string     `json:"-"`
	Role      string     `json:"role" example:"admin"`
}
