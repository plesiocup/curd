package handler

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/plesiocup/recommend/db"
	"gorm.io/gorm"
)

func Auth(c echo.Context) error {

	u := c.Get("user").(*jwt.Token)
	claims := u.Claims.(jwt.MapClaims)
	id := claims["id"].(float64)

	var user db.User
	if err := db.DB.Where("id = ?", id).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// return 404
			return c.JSON(http.StatusNotFound, echo.Map{
				"message": "User Not Found",
			})
		} else {
			// return 500
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"message": "Database Error: " + err.Error(),
			})
		}

	} else {
		// return 200
		return c.JSON(http.StatusOK, echo.Map{
			"id": id,
		})
	}
}
