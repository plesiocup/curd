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
	Id             uint   `gorm:"primaryKey"`
	Title          string `gorm:"not null"`
	Description    string `gorm:"not null"`
	Category       string `gorm:"not null"`
	Evaluation     uint   `gorm:"not null"`
	Playtime       string `gorm:"not null"`
	MovieURL       string `gorm:"not null"`
	ImageURL       string
	ReleaseYear    uint
	ClickedCount   uint
	EvaluatedCount uint
	SearchId       uint
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type UserbasedRecommend struct {
	Id         uint `gorm:"primaryKey"`
	UserId     uint
	MovieId    uint
	Evaluation uint
	Vector     int
}

// type ContentbasedRecommend struct {
// 	MovieId    uint `gorm:"primaryKey"`
// 	Evaluation uint
// 	Category   string
// }
