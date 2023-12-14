package db

import (
	"net/http"

	"gorm.io/gorm"
)

func CheckAdmin(id float64) (int, string) {

	var user User
	if err := DB.Where("id = ?", id).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// return 404
			return http.StatusNotFound, "User Not Found"
		} else {
			// return 500
			return http.StatusInternalServerError, "Database Error: " + err.Error()
		}
	}

	if user.Role != 0 {
		// return 403
		return http.StatusForbidden, "Permission Denied"
	} else {
		return 0, ""
	}
}
