package crawler

import (
	"auto-report/model"
	"errors"
	"github.com/devfeel/mapper"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"regexp"
	"strings"
)

var (
	LTURL = "http://xgsm.hitsz.edu.cn/zhxy-xgzs/xg_mobile/shsj/common"
	LOGINURL = "https://sso.hitsz.edu.cn:7002/cas/login?service=http://xgsm.hitsz.edu.cn/zhxy-xgzs/common/casLogin?params=L3hnX21vYmlsZS94c0hvbWU="
	POSTID = "http://xgsm.hitsz.edu.cn/zhxy-xgzs/xg_mobile/xs/csh"
	REPORT = "http://xgsm.hitsz.edu.cn/zhxy-xgzs/xg_mobile/xs/getYqxx"
	COMMITURL = "http://xgsm.hitsz.edu.cn/zhxy-xgzs/xg_mobile/xs/saveYqxx"
)

func init() {
	mapper.Register(&model.ReportData{})
	mapper.Register(&model.ModelData{})
}

func getLt(client http.Client) (string, error) {
	resp, err := client.Get(LTURL)
	if err != nil {
		log.Println(err)
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return "", err
	}
	template := regexp.MustCompile(`<input.*?type="hidden".*?value="(.*?)".*?/>`)
	lt := template.FindStringSubmatch(string(body))[1]
	return lt, nil
}

func Login(account, password string) (http.Client, error) {
	jar, _ := cookiejar.New(nil)
	var client = http.Client{
		Jar: jar,
	}
	lt, err := getLt(client)
	if err != nil {
		log.Println(err)
		return client, err
	}
	params := url.Values{
		"username":    {account},
		"password":    {password},
		"rememberMe":  {"on"},
		"lt":          {lt},
		"execution":   {"e1s1"},
		"_eventId":    {"submit"},
		"vc_username": {""},
		"vc_password": {""},
	}
	resp, err := client.Post(LOGINURL, "application/x-www-form-urlencoded", strings.NewReader(params.Encode()))
	if err != nil {
		log.Println(err)
		return client, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return client, err
	}
	bodyContent := string(body)
	if strings.Contains(bodyContent, "每日上报") {
		log.Println("登录成功！")
	} else {
		log.Println("登录失败！")
		return client, errors.New("login error")
	}
	return client, nil
}
