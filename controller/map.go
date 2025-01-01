package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/qw1281065768/trash_be/handler"
	"strconv"
	"strings"
)

// GetUserMaps - GET ALL
func GetUserMaps(c *gin.Context) {
	id := strings.TrimSpace(c.Query("uid"))
	userID, _ := strconv.ParseInt(id, 10, 64)
	fmt.Println(userID)
	resp := handler.GetTrashMapListByMainLevel(userID, 1)
	c.JSON(200, resp)
}

// GetMapInfo
func GetMapInfo(c *gin.Context) {
	id := strings.TrimSpace(c.Query("mapid"))
	mapID, _ := strconv.ParseInt(id, 10, 64)
	resp, err := handler.GetMapInfoByMapID(mapID)
	if err != nil {
		c.JSON(400, gin.H{})
	}
	c.JSON(200, resp)
}
