package db

import (
	"log"
	"os"

	"github.com/plesiocup/recommend/util"
)

func Migrate() {
	DB.Exec("DROP TABLE IF EXISTS users")
	DB.Exec("DROP TABLE IF EXISTS movies")

	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&Movie{})

	// test user 作成
	hashedPass, _ := util.HashPassword(os.Getenv("TESTUSER_PASSWORD"))
	user := User{
		Username: "user",
		Email:    "user@plesio.com",
		Password: hashedPass,
	}
	DB.Create(&user)

	log.Print("[INFO] DB Migrated!")
}
