package database

import (
	"fmt"
	"time"
)

type UserInfo struct {
	ID int64 `json:"id"`
	// 基础信息
	Level      int    `json:"level"`
	Name       string `json:"name"`
	Money      int    `json:"money"`
	Exp        int    `json:"exp"`
	CreateTime int64  `json:"create_time"`
	UpdateTime int64  `json:"update_time"`

	// 数值属性
	UserProperty UserProperty `json:"user_property"`

	// 装备佩戴
	Equipment Equipment `json:"equipment"`
}

type UserProperty struct {
	Power      float64 `json:"power"`      // 电力
	Efficiency float64 `json:"efficiency"` // 效率
	CritRate   float64 `json:"crit_rate"`  // 暴击
	Speed      float64 `json:"speed"`      // 速度
	Range      float64 `json:"range"`      // 范围
	Luck       float64 `json:"luck"`       // 幸运
	Weight     float64 `json:"weight"`     // 负重
}

// Equipment 定义装备结构体
type Equipment struct {
	Vehicles   string `json:"vehicles"`    // 载具
	Gloves     string `json:"gloves"`      // 手套
	PickupTool string `json:"pickup_tool"` // 拾取器
	Backpack   string `json:"backpack"`    // 背包
	Cleaner    string `json:"cleaner"`     // 清洁器
	Gem        string `json:"gem"`         // 宝石
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
