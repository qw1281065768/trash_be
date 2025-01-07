package database

import (
	"fmt"
	"github.com/qw1281065768/trash_be/model"
	"gorm.io/gorm"
	"time"
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

// 根据 user_id查询
func GetUserItemRelaByUserIDALL(userID int64) ([]UserItemRela, error) {
	db := GetDB()
	uirs := []UserItemRela{}
	err := db.Where("user_id = ?", userID).Find(&uirs).Error
	if err != nil {
		fmt.Println("error :", err)
		return nil, err
	}
	return uirs, nil
}

// 根据 user_id 和 item_type 查询
func GetUserItemRelaByUserIDANDType(userID int64, itemType int8) ([]UserItemRela, error) {
	db := GetDB()
	uirs := []UserItemRela{}
	err := db.Where("user_id = ? AND item_type = ?", userID, itemType).Find(&uirs).Error
	if err != nil {
		fmt.Println("error :", err)
		return nil, err
	}
	return uirs, nil
}

// AddItems 更新或添加物品到用户背包
func AddItems(userID int64, items map[int64]int) error {
	db := GetDB()
	for itemID, quantity := range items {
		var userItem UserItemRela
		// 查找用户-物品记录
		result := db.First(&userItem, "user_id = ? AND item_id = ?", userID, itemID)

		if result.Error != nil {
			if result.Error == gorm.ErrRecordNotFound {
				// 如果没有记录，创建新记录
				userItem = UserItemRela{
					UserID:     userID,
					ItemID:     itemID,
					ItemType:   int8(model.GlobalItemMap[itemID].Type), // 假设 item_type 设置为 1，您可以根据需要调整
					CreateTime: time.Now().Unix(),
					UpdateTime: time.Now().Unix(),
					Count:      quantity,
				}
				if err := db.Create(&userItem).Error; err != nil {
					return err
				}
			} else {
				return result.Error
			}
		} else {
			// 如果已有记录，更新数量和更新时间
			userItem.Count += quantity
			userItem.UpdateTime = time.Now().Unix()
			if err := db.Save(&userItem).Error; err != nil {
				return err
			}
		}
	}
	return nil
}

func UpdateUserItemCount(userID, itemID int64, count int) error {
	db := GetDB()
	var userItem UserItemRela
	// 查找用户-物品记录
	result := db.First(&userItem, "user_id = ? AND item_id = ?", userID, itemID)
	if result.Error != nil {
		return fmt.Errorf("no more items")
	}
	if userItem.Count < count {
		return fmt.Errorf("no enough items")
	}

	userItem.Count -= count
	userItem.UpdateTime = time.Now().Unix()
	if err := db.Save(&userItem).Error; err != nil {
		return err
	}
	return nil
}
