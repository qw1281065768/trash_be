package model

type HangingUser struct {
	ID          int            // 用户ID
	Bag         map[string]int // 物品名称，数量
	BagLimit    int            // 背包的容量
	OwnDropRate float64        // 用户自身的爆率，倍数
	IsHanging   bool           // 是否挂机中
	StartTime   int64          // 开始挂机时间
	EndTime     int64          // 挂机结束时间
	HangingTime int64          // 挂机持续时间
	TimeLimit   int64          // 挂机最长时间
	MapItems    []ItemsFall    // 地图爆的物品
}
