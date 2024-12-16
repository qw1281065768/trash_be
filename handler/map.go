package handler

import (
	"fmt"
	"github.com/qw1281065768/trash_be/database"
	"github.com/qw1281065768/trash_be/model"
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
