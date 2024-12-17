package controller

import (
	"github.com/gin-gonic/gin"
	grenderer "github.com/pilinux/gorest/lib/renderer"
	"github.com/qw1281065768/trash_be/handler"
	"strings"
)

// QueryString - basic implementation
func StartHanging(c *gin.Context) {
	query := strings.TrimSpace(c.Query("uid"))
	if query == "" {
		c.JSON(400, gin.H{"msg": query})
		return
	}

	// TODO 新增地图id + 工具列表

	// 需要判断用户是否解锁了地图 以及 是否拥有相关工具

	handler.StartHangingHandler(query)
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
