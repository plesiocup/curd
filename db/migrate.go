package db

import (
	"log"
	"os"

	"github.com/plesiocup/recommend/util"
)

func Migrate() {
	DB.Exec("DROP TABLE IF EXISTS users")
	DB.Exec("DROP TABLE IF EXISTS movies")
	DB.Exec("DROP TABLE IF EXISTS userbasedrecommends")

	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&Movie{})
	DB.AutoMigrate(&UserbasedRecommend{})

	// test user 作成
	hashedPass, _ := util.HashPassword(os.Getenv("TESTUSER_PASSWORD"))
	user := User{
		Username: "admin",
		Email:    "admin@plesio.com",
		Password: hashedPass,
		Role:     0, // admin
	}
	DB.Create(&user)

	log.Print("[INFO] DB Migrated!")
}
