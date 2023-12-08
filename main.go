package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/plesiocup/recommend/db"
	"github.com/plesiocup/recommend/handler"
)

func main() {
	godotenv.Load(".env")
	e := echo.New()
	db.Connect()
	// db.Migrate()

	e.POST("/users", handler.CreateUser)

	e.Logger.Fatal(e.Start(":8080"))
}
