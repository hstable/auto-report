package main

import (
	"auto-report/config"
	"auto-report/crawler"
	"log"
)

func main()  {
	info := config.GetConfig()
	isSuccess, err := crawler.Report(info.UserName, info.Password)
	if err != nil {
		log.Fatal("上报失败！")
		panic(err)
	}
	if isSuccess {
		// 发送邮件
	} else {
		log.Println("上报失败！")
	}
}
