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
