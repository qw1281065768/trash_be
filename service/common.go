package service

import (
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mediocregopher/radix/v4"
	log "github.com/sirupsen/logrus"

	"github.com/k3a/html2text"
	"github.com/pilinux/gorest/config"
	"github.com/pilinux/gorest/database"
	"github.com/pilinux/gorest/database/model"
	"github.com/pilinux/gorest/lib"
	"github.com/pilinux/gorest/lib/middleware"
	"github.com/pilinux/libgo/timestring"
	"gopkg.in/gomail.v2"
)

type EmailData struct {
	URL       string
	FirstName string
	Subject   string
}

// GetClaims - get JWT custom claims
func GetClaims(c *gin.Context) middleware.MyCustomClaims {
	// get claims
	claims := middleware.MyCustomClaims{
		AuthID:  c.GetUint64("authID"),
		Email:   c.GetString("email"),
		Role:    c.GetString("role"),
		Scope:   c.GetString("scope"),
		TwoFA:   c.GetString("tfa"),
		SiteLan: c.GetString("siteLan"),
		Custom1: c.GetString("custom1"),
		Custom2: c.GetString("custom2"),
	}

	return claims
}

// ValidateUserID - check whether authID or email is missing
func ValidateUserID(authID uint64, email string) bool {
	email = strings.TrimSpace(email)
	return authID != 0 && email != ""
}

// Validate2FA validates user-provided OTP
func Validate2FA(encryptedMessage []byte, issuer string, userInput string) ([]byte, string, error) {
	configSecurity := config.GetConfig().Security
	otpByte, err := lib.ValidateTOTP(encryptedMessage, issuer, userInput)
	// client provided invalid OTP / internal error
	if err != nil {
		// client provided invalid OTP
		if len(otpByte) > 0 {
			return otpByte, configSecurity.TwoFA.Status.Invalid, err
		}

		// internal error
		return []byte{}, "", err
	}

	// validated
	return otpByte, configSecurity.TwoFA.Status.Verified, nil
}

// DelMem2FA - delete secrets from memory
func DelMem2FA(authID uint64) {
	delete(model.InMemorySecret2FA, authID)
}

// SendEmail sends a verification/password recovery email if
// - required by the application
// - an external email service is configured
// - a redis database is configured
func SendEmail(email string, emailType int) bool {
	// send email if required by the application
	appConfig := config.GetConfig()

	// is external email service activated
	if appConfig.EmailConf.Activate != config.Activated {
		fmt.Println("email not active")
		return false
	}

	// is verification/password recovery email required
	doSendEmail := false
	if appConfig.Security.VerifyEmail && emailType == model.EmailTypeVerification {
		//fmt.Println("email type error1")
		//fmt.Println(emailType)
		doSendEmail = true
	}
	if appConfig.Security.RecoverPass && emailType == model.EmailTypePassRecovery {
		//fmt.Println("email type error2")
		doSendEmail = true
	}
	if !doSendEmail {
		fmt.Println("email type error3")
		return false
	}

	// is redis database activated
	if appConfig.Database.REDIS.Activate != config.Activated {
		fmt.Println("redis error")
		return false
	}

	data := struct {
		key   string
		value string
	}{}
	var keyTTL uint64
	var emailTag string
	var code uint64

	// generate verification/password recovery code
	if emailType == model.EmailTypeVerification {
		code = lib.SecureRandomNumber(appConfig.EmailConf.EmailVerificationCodeLength)
		data.key = model.EmailVerificationKeyPrefix + strconv.FormatUint(code, 10)
		keyTTL = appConfig.EmailConf.EmailVerifyValidityPeriod
		emailTag = appConfig.EmailConf.EmailVerificationTag
	}
	if emailType == model.EmailTypePassRecovery {
		code = lib.SecureRandomNumber(appConfig.EmailConf.PasswordRecoverCodeLength)
		data.key = model.PasswordRecoveryKeyPrefix + strconv.FormatUint(code, 10)
		keyTTL = appConfig.EmailConf.PassRecoverValidityPeriod
		emailTag = appConfig.EmailConf.PasswordRecoverTag
	}
	data.value = email

	// save in redis with expiry time
	client := *database.GetRedis()
	redisConnTTL := appConfig.Database.REDIS.Conn.ConnTTL

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(redisConnTTL)*time.Second)
	defer cancel()

	// Set key in Redis
	r1 := ""
	if err := client.Do(ctx, radix.FlatCmd(&r1, "SET", data.key, data.value)); err != nil {
		log.WithError(err).Error("error code: 401")
	}
	if r1 != "OK" {
		log.Error("error code: 402")
	}

	// Set expiry time
	r2 := 0
	if err := client.Do(ctx, radix.FlatCmd(&r2, "EXPIRE", data.key, keyTTL)); err != nil {
		log.WithError(err).Error("error code: 403")
	}
	if r2 != 1 {
		log.Error("error code: 404")
	}

	// check which email service
	// for Postmark
	if appConfig.EmailConf.Provider == "postmark" {
		htmlModel := lib.HTMLModel(lib.StrArrHTMLModel(appConfig.EmailConf.HTMLModel))
		htmlModel["secret_code"] = code
		htmlModel["email_validity_period"] = timestring.HourMinuteSecond(keyTTL)

		params := PostmarkParams{}
		params.ServerToken = appConfig.EmailConf.APIToken

		if emailType == model.EmailTypeVerification {
			params.TemplateID = appConfig.EmailConf.EmailVerificationTemplateID
		}

		if emailType == model.EmailTypePassRecovery {
			params.TemplateID = appConfig.EmailConf.PasswordRecoverTemplateID
		}

		params.From = appConfig.EmailConf.AddrFrom
		params.To = email
		params.Tag = emailTag
		params.TrackOpens = appConfig.EmailConf.TrackOpens
		params.TrackLinks = appConfig.EmailConf.TrackLinks
		params.MessageStream = appConfig.EmailConf.DeliveryType
		params.HTMLModel = htmlModel

		// send the email
		res, err := Postmark(params)
		if err != nil {
			log.WithError(err).Error("error code: 405")
		}
		if res.Message != "OK" {
			log.Error(res)
		}
	} else {
		fmt.Println("开发发邮件啦")
		var body bytes.Buffer

		temp, err := ParseTemplateDir("/home/ec2-user/gorest/gorest/templates")
		if err != nil {
			fmt.Println(err, "Could not parse template")
			//log.Fatal("Could not parse template", err)
		}
		// 👇 Send Email
		emailData := &EmailData{
			URL:       strconv.FormatInt(int64(code), 10),
			FirstName: "",
			Subject:   "Your account verification code",
		}
		err = temp.ExecuteTemplate(&body, "resetPassword.html", &emailData)
		if err != nil {
			fmt.Println(err)
		}
		// Sender data.
		from := "ai_gpt2023@163.com"
		smtpPass := "XGRWKLMDWWEFLIAI"
		smtpUser := "ai_gpt2023@163.com"
		to := email
		smtpHost := "smtp.163.com"
		smtpPort := 465
		fmt.Println(email, code)
		m := gomail.NewMessage()
		m.SetHeader("From", from)
		m.SetHeader("To", to)
		m.SetHeader("Subject", "Your account verification code")
		m.SetBody("text/html", body.String())
		m.AddAlternative("text/plain", html2text.HTML2Text(body.String()))

		d := gomail.NewDialer(smtpHost, smtpPort, smtpUser, smtpPass)
		d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
		// Send Email
		fmt.Println("开发发邮件成功!")
		if err := d.DialAndSend(m); err != nil {
			fmt.Println("Could not send email: ", err)
		}
	}

	return true
}

// 👇 Email template parser

func ParseTemplateDir(dir string) (*template.Template, error) {
	var paths []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	//fmt.Println(paths)
	return template.ParseFiles(paths...)
}
