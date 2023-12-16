package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/plesiocup/recommend/db"
	"github.com/plesiocup/recommend/util"
)

// スクレイピングして得た情報をDBに登録

func CreateMovie(c echo.Context) error {
	type Body struct {
		Title          string  `json:"title"`
		Description    string  `json:"description"`
		Category       string  `json:"category"`
		Evaluation     float64 `json:"evaluation"`
		Playtime       string  `json:"playtime"`
		MovieURL       string  `json:"movie_url"`
		ImageURL       string  `json:"image_url"`
		ReleaseYear    uint    `json:"release_year"`
		EvaluatedCount uint    `json:"evaluated_count"`
	}

	obj := new(Body)

	if err := c.Bind(obj); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Json Format Error: " + err.Error(),
		})
	}

	if util.HasEmptyField(obj, "Title", "Description", "Category", "MovieURL") {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "missing request field",
		})
	}

	new := db.Movie{
		Title:          obj.Title,
		Description:    obj.Description,
		Category:       obj.Category,
		Evaluation:     obj.Evaluation,
		Playtime:       obj.Playtime,
		MovieURL:       obj.MovieURL,
		ImageURL:       obj.ImageURL,
		ReleaseYear:    obj.ReleaseYear,
		EvaluatedCount: obj.EvaluatedCount,
	}
	db.DB.Create(&new)
	return c.JSON(http.StatusCreated, echo.Map{
		"title":           new.Title,
		"description":     new.Description,
		"category":        new.Category,
		"evaluation":      new.Evaluation,
		"playtime":        new.Playtime,
		"movie_url":       new.MovieURL,
		"image_url":       new.ImageURL,
		"release_year":    new.ReleaseYear,
		"evaluated_count": new.EvaluatedCount,
	})

}
