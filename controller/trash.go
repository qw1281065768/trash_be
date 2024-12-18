package controller

import (
	"github.com/gin-gonic/gin"
	grenderer "github.com/pilinux/gorest/lib/renderer"
	"github.com/qw1281065768/trash_be/handler"
	"strconv"
	"strings"
)

// QueryString - basic implementation
func StartHanging(c *gin.Context) {
	uid := strings.TrimSpace(c.Query("uid"))
	if uid == "" {
		c.JSON(400, gin.H{"uid": uid})
		return
	}
	mapIDStr := strings.TrimSpace(c.Query("mapid"))
	if mapIDStr == "" {
		c.JSON(400, gin.H{"mapid": mapIDStr})
		return
	}
	toolListStr := strings.TrimSpace(c.Query("tools"))
	if toolListStr == "" {
		c.JSON(400, gin.H{"tools": uid})
		return
	}

	// TODO 需要判断用户是否解锁了地图 以及 是否拥有相关工具
	mapID, _ := strconv.ParseInt(mapIDStr, 10, 64)
	toolList := strings.Split(toolListStr, ",")

	handler.StartHangingHandler(uid, mapID, toolList)
	grenderer.Render(c, nil, 200)
}

// GetPosts - GET /posts
func StopHanging(c *gin.Context) {
	query := strings.TrimSpace(c.Query("uid"))
	if query == "" {
		c.JSON(200, gin.H{"msg": query})
		return
	}

	handler.StopHangingHandler(query)
	grenderer.Render(c, nil, 200)
	//resp, statusCode := handler.GetPosts()
	//grenderer.Render(c, resp, statusCode)
}

/*type CheckBagRequest struct {
	UID string `json:"uid" binding:"required"`
}*/

// 使用idl生成代码，还需要返回总时间
func CheckBag(c *gin.Context) {
	query := strings.TrimSpace(c.Query("uid"))
	if query == "" {
		c.JSON(200, gin.H{"msg": query})
		return
	}

	resp, _ := handler.CheckUserBag(query)

	grenderer.Render(c, resp, 200)
}

// 查询所有的挂机任务
func CheckALLHanging(c *gin.Context) {
	resp := handler.CheckALLHanging()
	grenderer.Render(c, resp, 200)
}
