package model

import "time"

type Question struct {
	ID        uint       `gorm:"primary_key"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-" sql:"index"`

	question string   `json:"question" binding:"required"`
	options  []string ` json:"options" binding:"required"`
	answer   string   `json:"-" binding:"required"`
}
