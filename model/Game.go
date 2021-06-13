package model

import "time"

type Game struct {
	ID        uint       `gorm:"primary_key"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-" sql:"index"`

	Question  []Question ` gorm:"many2many:game_questions" json:"questions"`
	Score   uint   `json:"score"`
	Player []Users ` gorm:"many2many" json:"player"`
}
