package model

import (
	"fmt"
	"github.com/xuri/excelize/v2"
)

/*type Item struct {
	Name        string
	Probability float64 // 地图物品的基础爆率
}*/

type Item struct {
	ID             int64  `json:"id"`
	Name           string `json:"name"`
	Type           int    `json:"type"`
	TypeName       string `json:"type_name"`
	Property       string `json:"property"`
	PropertyName   string `json:"property_name"`
	DuringTime     int64  `json:"during_time"`
	ServLimitCount int64  `json:"servLimitCount"`
	Desc           string `json:"desc"`
	OriImgUrl      string `json:"ori_img_url"`
	Source         int    `json:"source"`
	Level          int    `json:"level"`
	Material       string `json:"material"`
	Weight         int    `json:"weight"`
	Exp            int    `json:"exp"`
	Price          int    `json:"price"`
}

var GlobalItemMap map[int64]Item

func init() {
	GlobalItemMap = make(map[int64]Item)

	// 2. 打开 Excel 文件
	file, err := excelize.OpenFile("./files/items.xlsx")
	if err != nil {
		fmt.Printf("Error opening Excel file: %v", err)
	}

	// 3. 获取表单数据
	sheetName := file.GetSheetName(0) // 获取第一个工作表
	rows, err := file.GetRows(sheetName)
	if err != nil {
		fmt.Printf("Error getting rows: %v", err)
	}

	// 4. 遍历 Excel 行并处理每一行数据
	for i, row := range rows {
		// 忽略表头
		if i == 0 {
			continue
		}
		// 解析每一行数据
		item := Item{
			ID:             parseInt64(row[0]),
			Name:           row[1],
			TypeName:       row[2],
			Type:           parseInt(row[3]),
			Property:       row[4],
			PropertyName:   row[5],
			DuringTime:     parseInt64(row[6]),
			ServLimitCount: parseInt64(row[7]),
			Desc:           row[8],
			OriImgUrl:      row[9],
			Source:         parseInt(row[10]),
			Level:          parseInt(row[11]),
			Material:       row[12],
			Weight:         parseInt(row[13]),
			Exp:            parseInt(row[14]),
			Price:          parseInt(row[15]),
		}
		GlobalItemMap[parseInt64(row[0])] = item
	}
}

func PrintMap() {
	fmt.Println(GlobalItemMap)
}

// 辅助函数：将字符串转换为整数
func parseInt(s string) int {
	var result int
	fmt.Sscanf(s, "%d", &result)
	return result
}

// 辅助函数：将字符串转换为 int64
func parseInt64(s string) int64 {
	var result int64
	fmt.Sscanf(s, "%d", &result)
	return result
}
