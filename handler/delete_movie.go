package handler

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/plesiocup/recommend/db"
	"gorm.io/gorm"
)

// 映画情報の削除（admin）
func DeleteMovie(c echo.Context) error {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userid := claims["id"].(float64)

	res, msg := db.CheckAdmin(userid)
	if res != 0 {
		return c.JSON(res, echo.Map{
			"message": msg,
		})
	}

	id := c.Param("id")
	var movie db.Movie
	if err := db.DB.Where("id = ?", id).First(&movie).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// return 404
			return c.JSON(http.StatusNotFound, echo.Map{
				"message": "Movie Not Found",
			})

		} else {
			// return 500
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"message": "Database Error: " + err.Error(),
			})
		}
	}

	db.DB.Delete(&db.Movie{}, id)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Deletion Successful",
	})
}
