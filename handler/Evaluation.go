package handler

import (
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/plesiocup/recommend/db"
)

func updateEvaluate(c echo.Context) error {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userid := claims["id"].(float64)

	movieidstr := c.Param("movieid")
	movieid, err := strconv.ParseUint(movieidstr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid movieid format",
		})
	}
	var movie db.Movie
	if err := db.DB.Where("id = ?", movieid).First(&movie).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Database Error: " + err.Error(),
		})
	} else {
		movie.Evaluation += 1
		movie.ClickedCount += 1

		db.DB.Save(&movie)
	}
	// 再計算＆更新
	newEvaluation := (movie.Evaluation*movie.EvaluatedCount + movie.ClickedCount) / (movie.EvaluatedCount + movie.ClickedCount)
	movie.Evaluation = newEvaluation
	db.DB.Save(&movie)
	// userbasedrecommendのcreate(userid,movieid,evaluation)
	new := db.UserbasedRecommend{
		UserId:     uint(userid),
		MovieId:    uint(movieid),
		Evaluation: newEvaluation,
	}
	db.DB.Create(&new)
	return c.JSON(http.StatusCreated, echo.Map{
		"UserId":     new.UserId,
		"MovieId":    new.MovieId,
		"Evaluation": new.Evaluation,
	})
	// ここからpythonをトリガー

}
