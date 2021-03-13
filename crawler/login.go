package crawler

import (
	"auto-report/model"
	"encoding/json"
	"errors"
	"fmt"
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
	// 获取每日上报 id
	req, err := http.NewRequest("POST", POSTID, nil)
	if err != nil {
		return err
	}
	resp, err = client.Do(req)
	if err != nil {
		log.Println(err)
		return err
	}
	body, _ = ioutil.ReadAll(resp.Body)
	var commitResult model.CommitResult
	err = json.Unmarshal(body, &commitResult)
	if err != nil {
		log.Println(err)
		return err
	}
	id := commitResult.Module
	// 点击 “新增“，获取默认信息
	reportId, err := json.Marshal(model.ID{id})
	if err != nil {
		log.Println(err)
		return err
	}
	params = url.Values{
		"info": {string(reportId)},
	}
	resp, err = client.Post(REPORT, "application/x-www-form-urlencoded; charset=UTF-8", strings.NewReader(params.Encode()))
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
	content := resultData.Module.Data[0]
	commitData := commit(&content)
	log.Println(commitData)
	// 提交
	cd, err := json.Marshal(commitData)
	if err != nil {
		log.Println(err)
		return err
	}
	params = url.Values{
		"info": {string(cd)},
	}
	resp, err = client.Post(COMMITURL, "application/x-www-form-urlencoded; charset=UTF-8", strings.NewReader(params.Encode()))
	if err != nil {
		log.Println(err)
		return err
	}
	body, _ = ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &commitResult)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println(commitResult)
	return nil
}

/**
利用默认的信息构造提交数据
*/
func commit(reportData *model.ReportData) model.CommitData {
	modelData := &model.ModelData{}
	mapper.Mapper(reportData, modelData)
	commitData := model.CommitData{
		Model: *modelData,
	}
	return commitData
}