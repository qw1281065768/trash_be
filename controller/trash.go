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
