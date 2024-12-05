package handler

import (
	"fmt"
	"github.com/qw1281065768/trash_be/model"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"
)

// 挂机玩法，涉及到的表和数据
var (
	users    = make(map[int]*model.User) // 存储多个用户
	mu       sync.Mutex                  // 锁以保护用户状态
	mapItems = []model.Item{             // 示例地图物品
		{"Item1", 0.3},  // 四等奖
		{"Item2", 0.2},  // 三等奖
		{"Item3", 0.05}, // 二等奖
		{"Item4", 0.05}, // 一等奖
	}
)

func InitUser(id int, ownDropRate float64) {
	mu.Lock()
	users[id] = &model.User{
		ID:          id,
		BagLimit:    10,
		OwnDropRate: ownDropRate,
		Bag:         make(map[string]int),
	}
	mu.Unlock()
}

func StartHangingHandler(UID string) {
	userID, _ := strconv.ParseInt(UID, 10, 64) // 这里根据请求获取用户ID
	//mu.Lock()
	fmt.Println(mapItems)
	user, exists := users[int(userID)]
	if !exists {
		fmt.Println("User not found", http.StatusNotFound)
		//mu.Unlock()
		return
	}
	if user.IsHanging {
		//mu.Unlock()
		fmt.Println("Already hanging")
		return
	}
	user.IsHanging = true
	//mu.Unlock()

	go hangUser(user)
	fmt.Printf("Hanging started for user %d\n", user.ID)
}

func StopHangingHandler(UID string) {
	userID, _ := strconv.ParseInt(UID, 10, 64)
	fmt.Printf("stopped for user %s\n", UID)
	//mu.Lock()
	user, exists := users[int(userID)]
	if !exists {
		fmt.Println("User not found", http.StatusNotFound)
		//mu.Unlock()
		return
	}
	user.IsHanging = false
	//mu.Unlock()
	fmt.Printf("Hanging stopped for user %d\n", user.ID)
}

func hangUser(user *model.User) {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()
	fmt.Printf("Hanging started for user %d\n", user.ID)
	for {
		//mu.Lock()
		if !user.IsHanging || len(user.Bag) >= user.BagLimit {
			//mu.Unlock()
			break
		}
		//mu.Unlock()

		searchItems(user)
		<-ticker.C
	}
	fmt.Printf("Hanging stopped for user %d\n", user.ID)
}

// 一次物品搜寻，爆率初始化为1（实际上就是抽奖次数，2倍的话就是抽两次）
func searchItems(user *model.User) {
	foundItems := make(map[string]int) // 存储捡到的物品及数量

	// 每次捡取物品的逻辑
	for _, item := range mapItems {
		if rand.Float64() < item.Probability*user.OwnDropRate { // 根据概率决定是否捡取
			foundItems[item.Name]++ // 增加捡到的物品数量
		}
	}

	mu.Lock()
	totalItems := 0 // 计算背包中物品的总数量
	for _, count := range user.Bag {
		totalItems += count
	}

	for itemName, count := range foundItems {
		if totalItems < user.BagLimit { // 检查总数量是否超过背包限制
			availableSpace := user.BagLimit - totalItems // 可用空间
			if count > availableSpace {
				count = availableSpace // 限制捡取数量
			}
			user.Bag[itemName] += count // 更新背包中的物品数量
			totalItems += count         // 更新背包中物品的总数量
			fmt.Printf("User %d got item: %s (x%d) | Bag: %v\n", user.ID, itemName, count, user.Bag)
		}
	}
	mu.Unlock()
}
