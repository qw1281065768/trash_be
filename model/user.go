package model

type User struct {
	ID          int
	Bag         map[string]int // 物品名称，数量
	BagLimit    int
	OwnDropRate float64 // 用户自身的爆率，倍数
	IsHanging   bool
}
