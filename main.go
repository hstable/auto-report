package main

import (
	"auto-report/config"
	"auto-report/crawler"
	"auto-report/email"
	"log"
)

func main() {
	info := config.GetConfig()
	isSuccess, err := crawler.Report(info.UserName, info.Password)
	if err != nil {
		log.Fatal("上报失败！")
		panic(err)
	}
	if isSuccess {
		log.Println("上报成功！")
	}
	if info.Email != "" {
		// 发送邮件
		err = email.SendEmail(info.Email, "上报成功!")
		if err != nil {
			log.Fatal(err)
		} else {
			log.Println("Sent email already!")
		}
	} else {
		log.Println("Unable to send email!")
	}
}
