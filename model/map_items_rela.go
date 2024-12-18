package model

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"strconv"
)

type ItemsFall struct {
	ItemID      int64   `json:"itemID"`
	ItemName    string  `json:"itemName"`
	Probability float64 `json:"probability"`
}

var MapItemsFall map[int64][]ItemsFall

func init() {
	MapItemsFall = make(map[int64][]ItemsFall)

	// 2. 打开 Excel 文件
	file, err := excelize.OpenFile("./files/map_item_rela.xlsx")
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
		if i == 0 || len(row) < 5 {
			continue
		}
		// 解析每一行数据
		mapID, _ := strconv.ParseInt(row[1], 10, 64)
		itemID, _ := strconv.ParseInt(row[2], 10, 64)
		prob, _ := strconv.ParseFloat(row[4], 64)
		ItemName := row[3]
		prob = prob / 100.0

		trashMap := ItemsFall{
			ItemID:      itemID,
			ItemName:    ItemName,
			Probability: prob,
		}

		MapItemsFall[mapID] = append(MapItemsFall[mapID], trashMap)
	}
}
