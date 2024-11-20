package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/pilinux/gorest/database/model"
	"github.com/pilinux/gorest/handler"
	"github.com/pilinux/gorest/lib/renderer"
)

// VerifyEmail - verify email address
func VerifyEmail(c *gin.Context) {
	payload := model.AuthPayload{}
	if err := c.ShouldBindJSON(&payload); err != nil {
		renderer.Render(c, gin.H{"message": err.Error()}, http.StatusBadRequest)
		return
	}

	resp, statusCode := handler.VerifyEmailV2(payload)

	renderer.Render(c, resp, statusCode)
}

// CreateVerificationEmail issues new verification code upon request
func CreateVerificationEmail(c *gin.Context) {
	payload := model.AuthPayload{}
	if err := c.ShouldBindJSON(&payload); err != nil {
		renderer.Render(c, gin.H{"message": err.Error()}, http.StatusBadRequest)
		return
	}

	resp, statusCode := handler.CreateVerificationEmail(payload)

	renderer.Render(c, resp, statusCode)
}

// SendVerificationEmail issues new verification code upon request
func SendVerificationEmail(c *gin.Context) {
	payload := model.AuthPayload{}
	if err := c.ShouldBindJSON(&payload); err != nil {
		renderer.Render(c, gin.H{"message": err.Error()}, http.StatusBadRequest)
		return
	}

	resp, statusCode := handler.CreateVerificationEmailV2(payload)

	renderer.Render(c, resp, statusCode)
}
