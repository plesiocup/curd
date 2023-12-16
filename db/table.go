package db

import "time"

type User struct {
	Id        uint   `gorm:"primaryKey"`
	Username  string `gorm:"not null"`
	Email     string `gorm:"not null;unique"`
	Password  string `gorm:"not null"`
	Role      uint   `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Movie struct {
	Id             uint    `gorm:"primaryKey"`
	Title          string  `gorm:"not null"`
	Description    string  `gorm:"not null"`
	Category       string  `gorm:"not null"`
	Evaluation     float64 `gorm:"not null"`
	Playtime       string  `gorm:"not null"`
	MovieURL       string  `gorm:"not null"`
	ImageURL       string
	ReleaseYear    uint
	EvaluatedCount uint
	SearchId       uint
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type UserbasedRecommend struct {
	Id         uint `gorm:"primaryKey"`
	UserId     uint
	MovieId    uint
	Evaluation float64
	Vector     int
}
