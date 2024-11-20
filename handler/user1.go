// Package handler ...
package handler

import (
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"

	gdatabase "github.com/pilinux/gorest/database"
	gmodel "github.com/pilinux/gorest/database/model"

	"github.com/pilinux/gorest/database/model"
)

// UpdateUser handles jobs for controller.UpdateUser
func UpdateUser1(email string, num int) (httpResponse gmodel.HTTPResponse, httpStatusCode int) {
	db := gdatabase.GetDB()
	userFinal := model.User1{}

	nowTime := time.Now()

	// does the user have an existing profile
	if err := db.Where("email = ?", email).First(&userFinal).Error; err != nil {

		// 用户不存在，新建用户，确保都是有vip到期时间的
		userFinal.CreatedAt = nowTime
		userFinal.UpdatedAt = nowTime
		userFinal.Vipdate = nowTime
		userFinal.Email = email

		tx := db.Begin()
		if err := tx.Create(&userFinal).Error; err != nil {
			// 创建失败
			tx.Rollback()
			log.WithError(err).Error("error code: 1111")
			httpResponse.Message = "internal server error"
			httpStatusCode = http.StatusInternalServerError
			return
		}
		tx.Commit()
	}

	if userFinal.Vipdate.Unix() < nowTime.Unix() {
		// 如果用户会员已到期，直接更新当前的时间
		userFinal.Vipdate = nowTime.Add(time.Hour * 24 * time.Duration(num))
	} else {
		// 未到期，加上现有的时间
		userFinal.Vipdate = userFinal.Vipdate.Add(time.Hour * 24 * time.Duration(num))
	}
	// user must not be able to manipulate all fields
	userFinal.UpdatedAt = nowTime

	tx := db.Begin()
	if err := tx.Save(&userFinal).Error; err != nil {
		tx.Rollback()
		log.WithError(err).Error("error code: 1121")
		httpResponse.Message = "internal server error"
		httpStatusCode = http.StatusInternalServerError
		return
	}
	tx.Commit()

	httpResponse.Message = userFinal
	httpStatusCode = http.StatusOK
	return
}

