package model

import (
	"fmt"
	"github.com/qw1281065768/trash_be/database"
)

type UserItemRela struct {
	ID         int64 `json:"id"`
	UserID     int64 `json:"user_id"`
	ItemID     int64 `json:"item_id"`
	ItemType   int8  `json:"item_type"`
	CreateTime int64 `json:"create_time"`
	UpdateTime int64 `json:"update_time"`
	Count      int   `json:"count"`
}

func (UserItemRela) TableName() string {
	return "user_item_rela" // 设置对应的表名
}

// 根据 user_id 和 item_type 查询
func GetUserItemRelaByUserIDAndType(userID int64, itemType int8) (*UserItemRela, error) {
	db := database.GetDB()
	uir := &UserItemRela{}
	err := db.Where("user_id = ?", userID).First(uir).Error
	if err != nil {
		fmt.Println("error :", err)
		return nil, err
	}
	fmt.Println(uir)
	return uir, nil
}
