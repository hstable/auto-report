package email

import (
	"errors"
	"gopkg.in/gomail.v2"
	"log"
	"strconv"
	"strings"
)

var (
	remoteInfo = map[string]string{
		"user": "dayreport@mzz.pub",
		"pass": "GoodGoodStudyDayDayReport!",
		"host": "smtp.mxhichina.com",
		"port": "465",
	}
)

func SendEmail(dest, content string) error {
	dest = strings.TrimSpace(dest)
	if dest == "" {
		return errors.New("邮件地址有误！")
	}
	port, err := strconv.Atoi(remoteInfo["port"])
	if err != nil {
		log.Println(err)
		return errors.New("邮件服务器端口解析有误！")
	}
	m := gomail.NewMessage()
	m.SetHeader("From", remoteInfo["user"])
	m.SetHeader("To", dest)
	m.SetHeader("Subject", "每日一报")
	m.SetBody("text/plain", content)
	d := gomail.NewDialer(remoteInfo["host"], port, remoteInfo["user"], remoteInfo["pass"])
	err = d.DialAndSend(m)
	return err
}
