package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pilinux/gorest/example/handler"

	grenderer "github.com/pilinux/gorest/lib/renderer"
)

// GetHobbies - GET /hobbies
func GetHobbies(c *gin.Context) {
	resp, statusCode := handler.GetHobbies()

	grenderer.Render(c, resp, statusCode)
}
