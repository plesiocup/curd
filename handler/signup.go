package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/plesiocup/recommend/db"
	"github.com/plesiocup/recommend/util"
	"gorm.io/gorm"
)

func CreateUser(c echo.Context) error {
	type Body struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Role     uint   `json:"role"`
		Password string `json:"password"`
	}

	obj := new(Body)
	// パース
	if err := c.Bind(obj); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Json Format Error: " + err.Error(),
		})
	}

	// 空のフィールドチェック
	if util.HasEmptyField(obj, "Username", "Email", "Password") {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "missing request field",
		})
	}
	// ユーザー登録
	var user db.User
	if err := db.DB.Where("email = ?", obj.Email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			hashedPass, err := util.HashPassword(obj.Password)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, echo.Map{
					"message": "Password Hashing Error",
				})
			}
			new := db.User{
				Username: obj.Username,
				Email:    obj.Email,
				Role:     obj.Role,
				Password: hashedPass,
			}
			db.DB.Create(&new)
			return c.JSON(http.StatusCreated, echo.Map{
				"id":         new.Id,
				"username":   new.Username,
				"email":      new.Email,
				"password":   new.Password,
				"role":       new.Role,
				"created_at": user.CreatedAt,
				"updated_at": user.UpdatedAt,
			})
		} else {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": "Database Error: " + err.Error(),
			})
		}
	} else {
		return c.JSON(http.StatusConflict, echo.Map{
			"message": "email conflict",
		})
	}
}
