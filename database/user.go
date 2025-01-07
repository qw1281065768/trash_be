package database

import (
	"fmt"
	"time"
)

type UserInfo struct {
	ID         int64  `json:"id"`
	Level      int    `json:"level"`
	Name       string `json:"name"`
	Money      int    `json:"money"`
	Exp        int    `json:"exp"`
	CreateTime int64  `json:"create_time"`
	UpdateTime int64  `json:"update_time"`
}

func (UserInfo) TableName() string {
	return "user" // 设置对应的表名
}

// 根据 user_id查询
func GetUserInfo(userID int64) (*UserInfo, error) {
	db := GetDB()
	user := &UserInfo{}
	err := db.Where("id = ?", userID).Find(user).Error
	if err != nil {
		fmt.Println("error :", err)
		return nil, err
	}
	return user, nil
}

func UpdateUser(user *UserInfo) error {
	db := GetDB()
	user.UpdateTime = time.Now().Unix()
	if err := db.Save(&user).Error; err != nil {
		return err
	}
	return nil
}
