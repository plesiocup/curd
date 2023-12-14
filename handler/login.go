package handler

import (
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/plesiocup/recommend/db"
	"github.com/plesiocup/recommend/util"
	"gorm.io/gorm"
)

func Login(c echo.Context) error {

	type Body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// parse
	obj := new(Body)
	if err := c.Bind(obj); err != nil {
		// return 400
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Json Format Error: " + err.Error(),
		})
	}

	if util.HasEmptyField(obj, "Username", "Email", "Password") {
		// return 400
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Missing Required Field",
		})
	}

	// ユーザーが存在するか
	var user db.User
	if err := db.DB.Where("email = ?", obj.Email).First(&user).Error; err != nil {
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
		if err := util.ComparePasswords(user.Password, obj.Password); err != nil {
			// return 401
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"message": "Invalid Password",
			})

		} else {
			// ペイロード作成
			claims := jwt.MapClaims{
				"id":  user.Id,
				"exp": time.Now().Add(time.Hour * 24).Unix(),
			}
			// トークン生成
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			// トークンに署名を付与
			tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
			if err != nil {
				return err
			}
			// return 200
			return c.JSON(http.StatusOK, echo.Map{
				"token": tokenString,
			})

		}
	}
}
