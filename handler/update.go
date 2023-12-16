package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/plesiocup/recommend/db"
)

// 情報のアプデ(バッチ内から)
func UpdateMovie(c echo.Context) error {
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
		SearchId       uint    `json:"search_id"`
	}

	id := c.Param("id")
	var movie db.Movie
	if err := db.DB.Where("id = ?", id).First(&movie).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Database Error: " + err.Error(),
		})

	} else {

		obj := new(Body)

		if err := c.Bind(obj); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": "Json Format Error: " + err.Error(),
			})
		}

		movie.Title = obj.Title
		movie.Description = obj.Description
		movie.Category = obj.Category
		movie.Evaluation = obj.Evaluation
		movie.Playtime = obj.Playtime
		movie.MovieURL = obj.MovieURL
		movie.ImageURL = obj.ImageURL
		movie.ReleaseYear = obj.ReleaseYear
		movie.EvaluatedCount = obj.EvaluatedCount
		movie.SearchId = obj.SearchId

		db.DB.Save(&movie)
		return c.JSON(http.StatusCreated, echo.Map{
			"title":        movie.Title,
			"description":  movie.Description,
			"category":     movie.Category,
			"evaluation":   movie.Evaluation,
			"playtime":     movie.Playtime,
			"movie_url":    movie.MovieURL,
			"image_url":    movie.ImageURL,
			"release_year": movie.ReleaseYear,
		})
	}
}
