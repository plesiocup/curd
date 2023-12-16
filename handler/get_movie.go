package handler

import (
	"net/http"
	"sort"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/plesiocup/recommend/db"
	"gorm.io/gorm"
)

// idで返すやつ
func GetMovie(c echo.Context) error {

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

	} else {
		// return 200
		return c.JSON(http.StatusCreated, echo.Map{
			"movie": movie,
		})

	}
}

// searchidで返すやつ(配列で受け取って配列で返す)
func GetSearchedMovie(c echo.Context) error {

	type Body struct {
		SearchId []uint `json:"search_id"`
	}

	obj := new(Body)

	if err := c.Bind(obj); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Json Format Error: " + err.Error(),
		})
	}

	// 映画IDの配列を元にレコードを検索
	var movies []db.Movie
	if err := db.DB.Find(&movies, obj.SearchId).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to search movies",
		})
	} else {
		// return 200
		return c.JSON(http.StatusOK, echo.Map{
			"movies": movies,
		})

	}
}

//
func GetUserRecommend(c echo.Context) error {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userid := claims["id"].(float64)

	var userrecommend []db.UserbasedRecommend
	if err := db.DB.Where("userid = ?", userid).Find(&userrecommend).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, echo.Map{
				"message": "Record Not Found",
			})
		} else {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": "Database Error" + err.Error(),
			})
		}
	} else {
		// vectorの上位10個のmovieIdを配列で返す
		sortedVectors := getSortedVector(userrecommend)
		res := getTopTens(sortedVectors)
		return c.JSON(http.StatusOK, echo.Map{
			"id": res,
		})
	}
}

// input: category, output: []Movie
func GetContentRecommend(c echo.Context) error {

	category := c.Param("category")

	var movies []db.Movie
	if err := db.DB.Where("category = ?", category).Find(&movies).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, echo.Map{
				"message": "Record Not Found",
			})
		} else {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": "Database Error" + err.Error(),
			})
		}
	} else {
		// そのカテゴリの映画一覧を配列で返す
		return c.JSON(http.StatusOK, echo.Map{
			"movies": movies,
		})
	}
}

// sort
func getSortedVector(userrec []db.UserbasedRecommend) []db.UserbasedRecommend {
	sort.Slice(userrec, func(i, j int) bool {
		return userrec[i].Vector > userrec[j].Vector
	})
	return userrec
}

// しぼりこみ
func getTopTens(userrec []db.UserbasedRecommend) []uint {
	var topTenMovie []uint
	for i := 0; i < len(userrec) && i < 10; i++ {
		topTenMovie = append(topTenMovie, userrec[i].MovieId)
	}
	return topTenMovie
}
