// Package controller contains all the controllers
// of the application
package controller

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"strconv"

	"github.com/gin-gonic/gin"

	grenderer "github.com/pilinux/gorest/lib/renderer"

	"github.com/pilinux/gorest/database/model"
	"github.com/pilinux/gorest/handler"
)

// GetUsers - GET /users
func GetUsers(c *gin.Context) {
	resp, statusCode := handler.GetUsers()

	grenderer.Render(c, resp, statusCode)
}

// GetUser - GET /users/:id
func GetUser(c *gin.Context) {
	id := strings.TrimSpace(c.Params.ByName("id"))

	resp, statusCode := handler.GetUser(id)

	if reflect.TypeOf(resp.Message).Kind() == reflect.String {
		grenderer.Render(c, resp, statusCode)
		return
	}

	grenderer.Render(c, resp.Message, statusCode)
}

// CreateUser - POST /users
func CreateUser(c *gin.Context) {
	userIDAuth := c.GetUint64("authID")
	user := model.User{}

	// bind JSON
	if err := c.ShouldBindJSON(&user); err != nil {
		grenderer.Render(c, gin.H{"message": err.Error()}, http.StatusBadRequest)
		return
	}

	resp, statusCode := handler.CreateUser(userIDAuth, user)

	if reflect.TypeOf(resp.Message).Kind() == reflect.String {
		grenderer.Render(c, resp, statusCode)
		return
	}

	grenderer.Render(c, resp.Message, statusCode)
}

// UpdateUser - PUT /users
func UpdateUser(c *gin.Context) {
	userIDAuth := c.GetUint64("authID")
	user := model.User{}

	// bind JSON
	if err := c.ShouldBindJSON(&user); err != nil {
		grenderer.Render(c, gin.H{"message": err.Error()}, http.StatusBadRequest)
		return
	}

	resp, statusCode := handler.UpdateUser(userIDAuth, user)

	if reflect.TypeOf(resp.Message).Kind() == reflect.String {
		grenderer.Render(c, resp, statusCode)
		return
	}

	grenderer.Render(c, resp.Message, statusCode)
}

// UpdateUser - PUT /users
func UpdateUser1(c *gin.Context) {
	email := c.Query("email")
	num,_ := strconv.Atoi(c.Query("num"))
	fmt.Println(email, num)
	user := model.User1{}
	// bind JSON
	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println("223")
		//grenderer.Render(c, gin.H{"message": err.Error()}, http.StatusBadRequest)
		//return
	}

	resp, statusCode := handler.UpdateUser1(email, num)

	if reflect.TypeOf(resp.Message).Kind() == reflect.String {
		fmt.Println("222343")
		grenderer.Render(c, resp, statusCode)
		return
	}

	grenderer.Render(c, resp.Message, statusCode)
}

// AddHobby - PUT /users/hobbies
func AddHobby(c *gin.Context) {
	userIDAuth := c.GetUint64("authID")
	hobby := model.Hobby{}

	// bind JSON
	if err := c.ShouldBindJSON(&hobby); err != nil {
		grenderer.Render(c, gin.H{"message": err.Error()}, http.StatusBadRequest)
		return
	}

	resp, statusCode := handler.AddHobby(userIDAuth, hobby)

	if reflect.TypeOf(resp.Message).Kind() == reflect.String {
		grenderer.Render(c, resp, statusCode)
		return
	}

	grenderer.Render(c, resp.Message, statusCode)
}
