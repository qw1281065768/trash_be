package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/qw1281065768/trash_be/handler"
	"strconv"

	"strings"
)

// GetUserItems - GET ALL
func GetUserItems(c *gin.Context) {
	id := strings.TrimSpace(c.Query("uid"))
	type1 := strings.TrimSpace(c.Query("type"))

	userID, _ := strconv.ParseInt(id, 10, 64)
	fmt.Println(userID)
	itemType, _ := strconv.ParseInt(type1, 10, 8)
	if itemType == 0 {
		itemType = 1
	}
	resp := handler.GetItemListALL(userID)
	c.JSON(200, resp)
	//grenderer.Render(c, resp, 0)
}

// SingleSellItems - GET 单个卖，以及多个卖
func SingleSellItems(c *gin.Context) {
	uid := strings.TrimSpace(c.PostForm("uid"))
	id := strings.TrimSpace(c.Query("item_id"))
	count := strings.TrimSpace(c.Query("count"))

	userID, _ := strconv.ParseInt(uid, 10, 64)
	itemID, _ := strconv.ParseInt(id, 10, 64)
	itemCount, _ := strconv.ParseInt(count, 10, 64)

	fmt.Println(userID)
	err := handler.SingleSellItem(userID, itemID, int(itemCount))
	if err != nil {
		c.JSON(400, nil)
	}
	c.JSON(200, nil)
}
