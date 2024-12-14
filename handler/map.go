package handler

import (
	"github.com/qw1281065768/trash_be/model"
)

func GetTrashMapList(userID int64) []model.TrashMap {

	mapList := make([]model.TrashMap, 0)
	/*resp, err := database.GetUserItemRelaByUserIDANDType(userID, int8(itemType))
	if err != nil {
		fmt.Println(err)
		return itemList
	}
	for _, v := range resp {
		itemInfo := model.GlobalItemMap[v.ItemID]
		itemInfo.OriImgUrl = "4001001"
		itemList = append(itemList, itemInfo)
		itemList = append(itemList, itemInfo)
	}*/

	// 获取用户的等级

	// 获取可以解锁的地图列表

	return mapList
}
