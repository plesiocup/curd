package handler

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/plesiocup/recommend/db"
	"gorm.io/gorm"
)

func GetUser(c echo.Context) error {
	user_t := c.Get("user").(*jwt.Token)
	claims := user_t.Claims.(jwt.MapClaims)
	userid := claims["id"].(float64)

	var user db.User
	if err := db.DB.Where("id = ?", userid).First(&user).Error; err != nil {

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
		return c.JSON(http.StatusCreated, echo.Map{
			"id":         user.Id,
			"username":   user.Username,
			"email":      user.Email,
			"role":       user.Role,
			"created_at": user.CreatedAt,
			"updated_at": user.UpdatedAt,
		})

	}
}
