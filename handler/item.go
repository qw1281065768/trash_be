package handler

import (
	"fmt"
	"github.com/qw1281065768/trash_be/database"
	"github.com/qw1281065768/trash_be/model"
)

func GetItemList(userID int64, itemType int) []model.Item {

	itemList := make([]model.Item, 0)
	resp, err := database.GetUserItemRelaByUserIDANDType(userID, int8(itemType))
	if err != nil {
		fmt.Println(err)
		return itemList
	}
	for _, v := range resp {
		itemInfo := model.GlobalItemMap[v.ItemID]
		itemInfo.OriImgUrl = "4001001"
		itemList = append(itemList, itemInfo)
		itemList = append(itemList, itemInfo)
	}

	return itemList
}

type ItemDetail struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	Type         int    `json:"type"`
	TypeName     string `json:"type_name"`
	Count        int    `json:"count"`
	Property     string `json:"property"`
	PropertyName string `json:"property_name"`
	Description  string `json:"desc"`
	Price        int    `json:"price"`
}

// GetItemListALL
func GetItemListALL(userID int64) []ItemDetail {
	itemList := make([]ItemDetail, 0)
	resp, err := database.GetUserItemRelaByUserIDALL(userID)
	if err != nil {
		fmt.Println(err)
		return itemList
	}
	for _, v := range resp {
		itemInfo := model.GlobalItemMap[v.ItemID]
		//itemInfo.OriImgUrl = "4001001"
		tmpItemInfo := ItemDetail{
			ID:           itemInfo.ID,
			Name:         itemInfo.Name,
			Type:         itemInfo.Type,
			TypeName:     itemInfo.TypeName,
			Property:     itemInfo.Property,
			PropertyName: itemInfo.PropertyName,
			Description:  itemInfo.Desc,
			Price:        itemInfo.Price,
			Count:        v.Count,
		}
		itemList = append(itemList, tmpItemInfo)
		//itemList = append(itemList, itemInfo)
	}

	return itemList
}

func SingleSellItem(userID int64, itemID int64, count int) error {
	// check 是否存在这么多物品
	itemList := GetItemListALL(userID)
	fmt.Println(itemList)
	exist := false
	sum := 0
	existCount := 0
	for _, v := range itemList {
		fmt.Println(v.ID, itemID)
		if v.ID == itemID {
			fmt.Println("111111")
			existCount = v.Count
			if v.Count >= count {
				exist = true
				// 计算总价格
				sum = v.Price * count
			} else {
				break
			}
		}
	}

	if !exist {
		return fmt.Errorf("count not enough: input : %d, exist : %d", count, existCount)
	}

	// 用户资产增加
	user, err := database.GetUserInfo(userID)
	if err != nil {
		return err
	}
	user.Money += sum
	err = database.UpdateUser(user)
	if err != nil {
		return err
	}

	// 物品数量扣除
	err = database.UpdateUserItemCount(userID, itemID, count)
	if err != nil {
		return err
	}

	// 返回结果
	return nil
}
