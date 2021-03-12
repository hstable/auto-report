package crawler

import (
	"auto-report/model"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"regexp"
	"strings"
)

var LTURL = "http://xgsm.hitsz.edu.cn/zhxy-xgzs/xg_mobile/shsj/common"
var LOGINURL = "https://sso.hitsz.edu.cn:7002/cas/login?service=http://xgsm.hitsz.edu.cn/zhxy-xgzs/common/casLogin?params=L3hnX21vYmlsZS94c0hvbWU="
var HISTORYREPORT = "http://xgsm.hitsz.edu.cn/zhxy-xgzs/xg_mobile/xs/getYqxxList"
var COMMITURL = "http://xgsm.hitsz.edu.cn/zhxy-xgzs/xg_mobile/xs/saveYqxx"

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
	fmt.Println(lt)
	return lt, nil
}

func Login(account, password string) error {
	jar, _ := cookiejar.New(nil)
	var client = http.Client{
		Jar: jar,
	}
	lt, err := getLt(client)
	if err != nil {
		log.Println(err)
		return err
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
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return err
	}
	bodyContent := string(body)
	if strings.Contains(bodyContent, "每日上报") {
		log.Println("登录成功！")
	} else {
		log.Println("登录失败！")
		return errors.New("login error")
	}
	// 获取每日上报历史信息
	req, err := http.NewRequest("POST", HISTORYREPORT, nil)
	if err != nil {
		return err
	}
	resp, err = client.Do(req)
	if err != nil {
		log.Println(err)
		return err
	}
	body, _ = ioutil.ReadAll(resp.Body)
	var resultData model.ResultData
	err = json.Unmarshal(body, &resultData)
	if err != nil {
		log.Println(err)
		return err
	}
	lastCommit := resultData.Module.Data[0]
	log.Println(lastCommit)
	// 利用上一次的疫情上报信息发送数据

	return nil
}