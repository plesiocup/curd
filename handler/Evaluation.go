package handler

import (
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/plesiocup/recommend/db"
)

func UpdateEvaluate(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userid := claims["id"].(float64)

	type Body struct {
		UserEval float64 `json:"user_eval"`
	}

	obj := new(Body)

	if err := c.Bind(obj); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Json Format Error: " + err.Error(),
		})
	}

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

		db.DB.Save(&movie)
	}
	// 再計算＆更新
	newEvaluation := (movie.Evaluation*float64(movie.EvaluatedCount) + obj.UserEval) / (float64(movie.EvaluatedCount) + 1)
	movie.Evaluation = newEvaluation
	movie.EvaluatedCount += 1
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
