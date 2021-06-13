package model

import "time"

type Options struct {
	ID        uint       `gorm:"primary_key,omitempty"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-" sql:"index"`

	Option     string `json:"option" binding:"required"`
	QuestionID uint   `json:"question_id" binding:"required"`
}