package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/plesiocup/recommend/db"
	"gorm.io/gorm"
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

		// 再計算＆更新
		movie.EvaluatedCount += 1
		newEvaluation := (movie.Evaluation*float64(movie.EvaluatedCount) + obj.UserEval) / (float64(movie.EvaluatedCount) + 1)
		movie.Evaluation = newEvaluation
		db.DB.Save(&movie)
		// userbasedrecommendのcreate(userid,movieid,evaluation)
		new := db.UserbasedRecommend{
			UserId:     uint(userid),
			MovieId:    uint(movieid),
			Evaluation: newEvaluation,
		}
		db.DB.Create(&new)
		// ここからpythonをトリガー

		var result map[string]interface{}
		query := `
		SELECT JSON_ARRAYAGG(JSON_OBJECT(
			'userId', user_id, 
			'movieId', movie_id,
			'evaluation', evaluation
		)) AS jsonData
		FROM userbased_recommends;
	`
		var jsonData string
		if err := db.DB.Raw(query).Scan(&jsonData).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"message": "Error executing SQL query",
			})
		}

		// FastAPIエンドポイントのURL

		apiURL := "https://python-recommend.azurewebsites.net/userbasedrecommend/" + strconv.Itoa(int(userid))

		jsonDataBytes := []byte(jsonData)
		resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(jsonDataBytes))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"message": "Error sending HTTP request:" + err.Error(),
			})
		}
		defer resp.Body.Close()
		// レスポンスの読み取り
		json.NewDecoder(resp.Body).Decode(&result)

		result = map[string]interface{}{
			"recommendations": []interface{}{1, 2},
		}
		moviesList, _ := getMovieList(c, result)

		res := updateUserRecommend(c, moviesList, uint(userid))

		return c.JSON(http.StatusCreated, echo.Map{
			"massage": res,
		})
	}

}
func getMovieList(c echo.Context, result map[string]interface{}) ([]int, error) {
	// resultからrecommendationsを取り出す
	recommendationsInterface, ok := result["recommendations"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("Invalid recommendations format")
	}
	// recommendationsをint[]に変換
	var movieIds []int
	for _, id := range recommendationsInterface {
		if idInt, ok := id.(int); ok {
			movieIds = append(movieIds, int(idInt))
		} else {
			fmt.Printf("Invalid movie ID format: %v\n", id)
		}

	}
	return movieIds, nil

}

func updateUserRecommend(c echo.Context, movies []int, userId uint) error {
	var moviesid []uint

	for _, intValue := range movies {
		if intValue < 0 {
			return fmt.Errorf("Negative integer values are not supported: %d", intValue)
		}
		moviesid = append(moviesid, uint(intValue))
	}
	var user db.User
	if err := db.DB.Where("id = ?", userId).First(&user).Error; err != nil {
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
		for _, movieID := range moviesid {
			log.Print(movieID)
			query := fmt.Sprintf("INSERT INTO user_recommendations (user_id, movie_id) VALUES (%d, %d);", userId, movieID)
			if err := db.DB.Exec(query).Error; err != nil {
				return err
			}
		}

		return c.JSON(http.StatusCreated, echo.Map{
			"message": "Recommend Update",
		})
	}
}
