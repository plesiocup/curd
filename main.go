package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/plesiocup/recommend/db"
	"github.com/plesiocup/recommend/handler"
)

func main() {
	godotenv.Load(".env")
	e := echo.New()
	db.Connect()
	// db.Migrate()

	config := middleware.JWTConfig{
		SigningKey: []byte(os.Getenv("JWT_SECRET_KEY")),
		ParseTokenFunc: func(tokenString string, c echo.Context) (interface{}, error) {
			keyFunc := func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(os.Getenv("JWT_SECRET_KEY")), nil
			}

			token, err := jwt.Parse(tokenString, keyFunc)
			if err != nil {
				return nil, err
			}
			if !token.Valid {
				return nil, errors.New("invalid token")
			}
			return token, nil
		},
	}

	e.POST("/signup", handler.CreateUser)
	// e.GET("/user", handler.GetUser)

	e.POST("/login", handler.Login)

	e.POST("/movie", handler.CreateMovie)
	e.PUT("/movie/:id", handler.UpdateMovie)

	r := e.Group("/auth")
	r.Use(middleware.JWTWithConfig(config))

	r.GET("", handler.Auth)

	r.GET("/getMovie/:id", handler.GetMovie)
	r.GET("/getMovie", handler.GetSearchedMovie)
	r.GET("/userbasedRecommend", handler.GetUserRecommend)
	r.GET("/contentbasedRecommend/:category", handler.GetContentRecommend)

	r.PUT("/evaluate/:movieid", handler.UpdateEvaluate)
	// r.DELETE("/movie", handler.DeleteMovie)

	e.Logger.Fatal(e.Start(":8080"))
}
