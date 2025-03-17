package email

import (
	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
	"strconv"
)

func SendCode(aimEmail string, code string) error {
	var emailHost = viper.GetString("email.host")
	var username = viper.GetString("email.username")
	var password = viper.GetString("email.password")
	var port = viper.GetString("email.port")
	m := gomail.NewMessage()
	m.SetHeader("From", username)
	m.SetHeader("To", aimEmail)
	m.SetHeader("Subject", "验证码")
	m.SetBody("text/html", "您本次的验证码是："+code)
	var returnErr error
	if port, err := strconv.Atoi(port); err == nil {
		d := gomail.NewDialer(emailHost, port, username, password)
		if err := d.DialAndSend(m); err != nil {
			returnErr = err
		}
	} else {
		returnErr = err
	}
	return returnErr
}
