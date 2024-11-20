// Package controller contains all the controllers
// of the application
package controller

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"

	"github.com/pilinux/gorest/database/model"
	"github.com/pilinux/gorest/handler"
	"github.com/pilinux/gorest/lib/renderer"
)

// CreateUserAuth - POST /register
func CreateUserAuth(c *gin.Context) {
	auth := model.Auth{}

	// bind JSON
	if err := c.ShouldBindJSON(&auth); err != nil {
		renderer.Render(c, gin.H{"message": err.Error()}, http.StatusBadRequest)
		return
	}

	resp, statusCode := handler.CreateUserAuth(auth)

	if reflect.TypeOf(resp.Message).Kind() == reflect.String {
		renderer.Render(c, resp, statusCode)
		return
	}

	renderer.Render(c, resp.Message, statusCode)
}
