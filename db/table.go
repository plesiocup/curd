package db

import "time"

type User struct {
	Id        uint   `gorm:"primaryKey"`
	Username  string `gorm:"not null"`
	Email     string `gorm:"not null;unique"`
	Password  string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
type Movie struct {
	Id          uint   `gorm:"prmaryKey"`
	Title       string `gorm:"not null"`
	Description string `gorm:"not null"`
	Category    string `gorm:"not null"`
	Playtime    string `gorm:"not null"`
	Review      uint
	ReleaseYear uint
	Img         string
	AmazonURL   string `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
