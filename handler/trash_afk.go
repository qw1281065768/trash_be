package handler

import (
	"errors"
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
	userPool = make(map[int]*model.User) // 挂机池子，先用本地map存储，后续替换成分布式的

	mu       sync.Mutex // 锁以保护用户状态
	mapItems = []model.Item{ // 示例地图物品
		{"Item1", 0.3},  // 四等奖
		{"Item2", 0.2},  // 三等奖
		{"Item3", 0.05}, // 二等奖
		{"Item4", 0.05}, // 一等奖
	}
)

// 初始化用户数据，一般是在用户开始挂机的时候来使用
func InitUser(id int, ownDropRate float64) *model.User {
	mu.Lock()
	userPool[id] = &model.User{
		ID:          id,
		BagLimit:    10,
		TimeLimit:   300,
		OwnDropRate: ownDropRate,
		Bag:         make(map[string]int),
		StartTime:   time.Now().Unix(),
	}
	mu.Unlock()
	return userPool[id]
}

func StartHangingHandler(UID string) {
	userID, _ := strconv.ParseInt(UID, 10, 64) // 这里根据请求获取用户ID
	fmt.Println(mapItems)
	user, exists := userPool[int(userID)]
	if !exists {
		fmt.Println("User not found")
		user = InitUser(int(userID), 1)
	}
	if user.IsHanging {
		fmt.Println("Already hanging")
		return
	}
	go hangUser(user)
	fmt.Printf("Hanging started for user %d\n", user.ID)
}

func StopHangingHandler(UID string) {
	userID, _ := strconv.ParseInt(UID, 10, 64)
	fmt.Printf("stopped for user %s\n", UID)
	user, exists := userPool[int(userID)]
	if !exists {
		fmt.Println("User not found", http.StatusNotFound)
		return
	}
	shutdownHanging(user)
	fmt.Printf("Hanging stopped for user %d\n", user.ID)
}

func hangUser(user *model.User) {
	ticker := time.NewTicker(5 * time.Second)
	user.IsHanging = true
	defer ticker.Stop()
	fmt.Printf("Hanging started for user %d\n", user.ID)
	for {
		// 状态判断
		if !user.IsHanging {
			break
		}
		// 超时判断
		if time.Now().Unix()-user.StartTime > user.TimeLimit {
			shutdownHanging(user)
			break
		}
		searchItems(user)
		<-ticker.C
	}
	fmt.Printf("Hanging stopped for user %d\n", user.ID)
}

// 一次物品搜寻，爆率初始化为1（实际上就是抽奖次数，2倍的话就是抽两次）
func searchItems(user *model.User) {
	foundItems := make(map[string]int) // 存储捡到的物品及数量

	fmt.Printf("Searching items for user %d\n", user.ID)
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
			if count >= availableSpace {
				count = availableSpace // 限制捡取数量
				shutdownHanging(user)  // 终止挂机
			}
			user.Bag[itemName] += count // 更新背包中的物品数量
			totalItems += count         // 更新背包中物品的总数量
			fmt.Printf("User %d got item: %s (x%d) | Bag: %v\n", user.ID, itemName, count, user.Bag)
		}
	}
	mu.Unlock()
}

// CheckBagResponse 返回最近一次挂机的数据
type CheckBagResponse struct {
	HangingStartTime    int64          `json:"hanging_start_time"`     // 挂机开始时间，时间戳
	HangingStartTimeStr string         `json:"hanging_start_time_str"` // 挂机开始时间，日期+时分秒
	DuringTime          int64          `json:"during_time"`            // 挂机时长
	IsHanging           bool           `json:"is_hanging"`             // 是否挂机中
	BagLimit            int            `json:"bag_limit"`              // 背包容量
	UserID              string         `json:"user_id"`                // 用户id
	BagContent          map[string]int `json:"bag_content"`            // 背包内容，物品名称+数量
}

// CheckUserBag 查找用户上一次挂机的信息
func CheckUserBag(UID string) (*CheckBagResponse, error) {
	resp := &CheckBagResponse{
		UserID: UID,
	}
	userID, _ := strconv.ParseInt(UID, 10, 64)
	user, exists := userPool[int(userID)]
	if !exists {
		fmt.Println("User not found")
		return nil, errors.New("User not found")
	}
	resp.BagLimit = user.BagLimit
	resp.BagContent = user.Bag
	resp.HangingStartTime = user.StartTime
	resp.IsHanging = user.IsHanging

	if user.IsHanging {
		fmt.Println("Already hanging")
		resp.DuringTime = time.Now().Unix() - user.StartTime
	} else {
		resp.DuringTime = user.EndTime - user.StartTime
	}

	return resp, nil
}

// CheckALLHanging 查询整体的挂机
func CheckALLHanging() map[int]*model.User {
	for _, user := range userPool {
		if user.IsHanging {
			user.HangingTime = time.Now().Unix() - user.StartTime
		} else {
			user.HangingTime = user.EndTime - user.StartTime
		}
	}
	return userPool
}

func shutdownHanging(user *model.User) {
	user.IsHanging = false
	user.EndTime = time.Now().Unix()
}
