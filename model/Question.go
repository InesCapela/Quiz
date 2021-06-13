package model

import "time"

type Question struct {
	ID uint `gorm:"primary_key,omitempty"`

	Question string `json:"question" binding:"required"`
	Options   []Options `json:"options" binding:"required"`
	Answer string   `json:"answer,omitempty" binding:"required"`

	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-" sql:"index"`
}
