package email

import (
	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
	"log"
	"strconv"
)

func Send(aimEmail string) {
	log.Println("发送邮件给", aimEmail)
	var emailHost = viper.GetString("email.host")
	var username = viper.GetString("email.username")
	var password = viper.GetString("email.password")
	var port = viper.GetString("email.port")
	m := gomail.NewMessage()
	m.SetHeader("From", username)
	m.SetHeader("To", aimEmail)
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", "Hello <b>herzorf</b>!")
	if port, err := strconv.Atoi(port); err == nil {
		d := gomail.NewDialer(emailHost, port, username, password)
		if err := d.DialAndSend(m); err != nil {
			panic(err)
		}
	}

}
