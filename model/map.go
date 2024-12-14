package model

import (
	"fmt"
	"github.com/xuri/excelize/v2"
)

type TrashMap struct {
	ID         int64  `json:"id"`
	MainLevel  string `json:"mainLevel"`
	SubLevel   string `json:"subLevel"`
	BGImage    string `json:"bgImage"`
	UserLevel  int    `json:"userLevel"`
	ProdSpeed  int    `json:"prodSpeed"`
	TotalItems int    `json:"totalItems"`
	Desc       string `json:"desc"`
}

var GlobalTrashMap map[int64]TrashMap

func init() {
	GlobalTrashMap = make(map[int64]TrashMap)

	// 2. 打开 Excel 文件
	file, err := excelize.OpenFile("./files/maps.xlsx")
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
		trashMap := TrashMap{
			ID:         parseInt64(row[0]),
			MainLevel:  row[1],
			SubLevel:   row[2],
			BGImage:    row[3],
			UserLevel:  parseInt(row[4]),
			ProdSpeed:  parseInt(row[5]),
			TotalItems: parseInt(row[6]),
			Desc:       row[7],
		}
		GlobalTrashMap[parseInt64(row[0])] = trashMap
	}
}
