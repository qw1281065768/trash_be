package handler

import (
	"fmt"
	"github.com/qw1281065768/trash_be/database"
	"github.com/qw1281065768/trash_be/model"
	"sort"
)

// GetTrashMapListByMainLevel 根据主关卡获取已解锁的关卡名称
func GetTrashMapListByMainLevel(userID int64, mainLevel int) []model.TrashMap {

	mapList := make([]model.TrashMap, 0)
	user, err := database.GetUserInfo(userID)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	// 获取用户的等级
	fmt.Println(user.Level)

	// 获取可以解锁的地图列表

	return mapList
}

type MapInfo struct {
	Info      model.TrashMap
	ItemFalls []model.ItemsFall
}

func GetMapInfoByMapID(mapID int64) (*MapInfo, error) {
	resp := &MapInfo{}
	itemFalls := model.MapItemsFall[mapID]
	if len(itemFalls) == 0 {
		// 地图里没有掉落的物品
		return resp, fmt.Errorf("empty map items")
	}
	// 排序
	sort.Slice(itemFalls, func(i, j int) bool {
		return itemFalls[i].Probability > itemFalls[j].Probability // 根据 ID 排序
	})

	resp.ItemFalls = itemFalls
	info, ok := model.GlobalTrashMap[mapID]
	if !ok {
		// 地图里没有掉落的物品
		return resp, fmt.Errorf("empty map info")
	}
	resp.Info = info
	return resp, nil
}
