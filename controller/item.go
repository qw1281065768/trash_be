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
	resp := handler.GetItemList(userID, int(itemType))
	c.JSON(200, resp)
	//grenderer.Render(c, resp, 0)
}
