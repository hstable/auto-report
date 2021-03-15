package config

import (
	"flag"
	"log"
	"sync"
)

type Config struct {
	UserName string
	Password string
	Email    string
}

var config Config
var once sync.Once

func init() {
	flag.StringVar(&config.UserName, "u", "", "Your student ID, 学号, e.g., 20S051030")
	flag.StringVar(&config.Password, "p", "", "Your password,密码")
	flag.StringVar(&config.Email, "e", "", "Your email, 接收邮件的邮箱")
}

func parseParams() {
	flag.Parse()
	if config.UserName == "" || config.Password == "" {
		log.Fatal("Neither USERNAME nor PASSWORD can be empty! See help by command line flag -help.")
	}
}

func GetConfig() *Config {
	once.Do(parseParams)
	return &config
}
