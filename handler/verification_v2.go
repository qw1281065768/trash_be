package handler

import (
	"context"
	"errors"
	"github.com/mediocregopher/radix/v4"
	"github.com/pilinux/gorest/config"
	"github.com/pilinux/gorest/database"
	"github.com/pilinux/gorest/database/model"
	"github.com/pilinux/gorest/lib"
	"github.com/pilinux/gorest/service"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"time"
)

// CreateVerificationEmailV2 handles jobs for controller.CreateVerificationEmail  直接给邮箱发消息
func CreateVerificationEmailV2(payload model.AuthPayload) (httpResponse model.HTTPResponse, httpStatusCode int) {
	payload.Email = strings.TrimSpace(payload.Email)
	// 邮箱格式校验
	if !lib.ValidateEmail(payload.Email) {
		httpResponse.Message = "wrong email address"
		httpStatusCode = http.StatusBadRequest
		httpResponse.Code = model.DEFUALT_ERROR
		return
	}

	// 直接发邮件
	if !service.SendEmail(payload.Email, model.EmailTypeVerification) {
		httpResponse.Message = "failed to send verification email"
		httpStatusCode = http.StatusServiceUnavailable
		httpResponse.Code = model.DEFUALT_ERROR
		return
	}

	httpResponse.Message = "sent verification email"
	httpResponse.Code = model.SUCCESS
	httpStatusCode = http.StatusOK
	return
}

// VerifyEmail handles jobs for controller.VerifyEmail  没用到邮箱，直接去查db了，可以优化一个新方法
func VerifyEmailV2(payload model.AuthPayload) (httpResponse model.HTTPResponse, httpStatusCode int) {
	data := struct {
		key   string
		value string
	}{}

	// 直接把验证码加到redis里，不带邮箱
	data.key = model.EmailVerificationKeyPrefix + payload.VerificationCode

	// get redis client
	client := *database.GetRedis()
	rConnTTL := config.GetConfig().Database.REDIS.Conn.ConnTTL
	// 不影响主协程
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(rConnTTL)*time.Second)
	defer cancel()

	// is key available in redis
	result := 0
	if err := client.Do(ctx, radix.FlatCmd(&result, "EXISTS", data.key)); err != nil {
		log.WithError(err).Error("error code: 1061")
		httpResponse.Message = "internal server error"
		httpStatusCode = http.StatusInternalServerError
		httpResponse.Code = model.DEFUALT_ERROR
		return
	}

	// 不存在或者过期
	if result == 0 {
		httpResponse.Message = "wrong/expired verification code"
		httpStatusCode = http.StatusUnauthorized
		httpResponse.Code = model.DEFUALT_ERROR
		return
	}

	// find key in redis
	if err := client.Do(ctx, radix.FlatCmd(&data.value, "GET", data.key)); err != nil {
		log.WithError(err).Error("error code: 1062")
		httpResponse.Message = "internal server error"
		httpStatusCode = http.StatusInternalServerError
		httpResponse.Code = model.DEFUALT_ERROR
		return
	}

	// delete key from redis
	result = 0
	if err := client.Do(ctx, radix.FlatCmd(&result, "DEL", data.key)); err != nil {
		log.WithError(err).Error("error code: 1063")
	}
	if result == 0 {
		err := errors.New("failed to delete recovery key from redis")
		log.WithError(err).Error("error code: 1064")
	}

	// 校验邮箱是否一致
	if data.value != payload.Email {
		httpResponse.Message = "wrong verification code"
		httpStatusCode = http.StatusInternalServerError
		httpResponse.Code = model.DEFUALT_ERROR
		return
	}

	httpResponse.Message = "email successfully verified"
	httpResponse.Code = model.SUCCESS
	httpStatusCode = http.StatusOK
	return
}

