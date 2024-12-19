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
