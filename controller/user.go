// Package controller contains all the controllers
// of the application
package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/qw1281065768/trash_be/handler"
	"strconv"
	"strings"
)

func GetUserInfo(c *gin.Context) {
	uid := strings.TrimSpace(c.Query("uid"))
	userID, _ := strconv.ParseInt(uid, 10, 64)
	resp, err := handler.GetUserInfo(userID)
	if err != nil {
		c.JSON(400, nil)
	}
	c.JSON(200, resp)
}
