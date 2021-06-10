package model

import "time"

type Question struct {
	ID uint `gorm:"primary_key"`

	Question string `json:"question" binding:"required"`
	Options   []Options `json:"options" binding:"required" gorm:"foreign:QuestionID"`
	//Option []string `json:"options" binding:"required" gorm:"embedded"`
	Answer string   `json:"answer" binding:"required"`

	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-" sql:"index"`
}
