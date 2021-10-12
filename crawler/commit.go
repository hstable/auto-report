package crawler

import (
	"auto-report/model"
	"encoding/json"
	"fmt"
	"github.com/devfeel/mapper"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"
)

/**
自动上报
*/
func Report(username, password string) (bool, error) {
	// 登录学校网站
	client, err := login(username, password)
	if err != nil {
		log.Println("Report: login: " + err.Error())
		return false, err
	}
	// 获取每日上报 Id
	id, err := getId(client)
	if err != nil {
		log.Println("Report: getId: " + err.Error())
		return false, err
	}
	// 点击 “新增“，获取默认信息
	commitData, err := pressNewButton(client, id)
	if err != nil {
		log.Println("Report: pressNewButton: " + err.Error())
		return false, err
	}
	// 检查温度信息
	commitData = checkTemperature(commitData)
	// 提交
	result, err := commit(client, commitData)
	if err != nil {
		log.Println("Report: commit: " + err.Error())
		return false, err
	}
	return result.IsSuccess, nil
}

func checkTemperature(commitData model.CommitData) model.CommitData {
	if commitData.Model.Brzgtw == "" {
		rand.Seed(time.Now().Unix())
		temp := fmt.Sprintf("%.1f", 36 + rand.Float32())
		commitData.Model.Brzgtw = temp
	}
	return commitData
}

// 获取每日上报 Id
func getId(client http.Client) (string, error) {
	req, err := http.NewRequest("POST", POSTID, nil)
	if err != nil {
		return "", err
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("getId: client.Do: " + err.Error())
		return "", err
	}
	body, _ := ioutil.ReadAll(resp.Body)
	var commitResult model.CommitResult
	err = json.Unmarshal(body, &commitResult)
	if err != nil {
		log.Println("getId: json.Unmarshal: " + err.Error())
		return "", err
	}
	id := commitResult.Module
	return id, nil
}

/**
点击 “新增“，获取默认信息
*/
func pressNewButton(client http.Client, id string) (model.CommitData, error) {
	reportId, err := json.Marshal(model.ID{id})
	if err != nil {
		log.Println("pressNewButton: json.Marshal: " + err.Error())
		return model.CommitData{}, err
	}
	params := url.Values{
		"info": {string(reportId)},
	}
	resp, err := client.Post(REPORT, "application/x-www-form-urlencoded; charset=UTF-8", strings.NewReader(params.Encode()))
	if err != nil {
		log.Println(err)
		return model.CommitData{}, err
	}
	body, _ := ioutil.ReadAll(resp.Body)
	var resultData model.ResultData
	err = json.Unmarshal(body, &resultData)
	if err != nil {
		log.Println("pressNewButton: json.Unmarshal: " + err.Error())
		return model.CommitData{}, err
	}
	content := resultData.Module.Data[0]
	commitData := genCommitData(&content)
	return commitData, nil
}

/**
提交
*/
func commit(client http.Client, commitData model.CommitData) (model.CommitResult, error) {
	var commitResult model.CommitResult
	cd, err := json.Marshal(commitData)
	if err != nil {
		log.Println("commit: json.Marshal: " + err.Error())
		return commitResult, err
	}
	params := url.Values{
		"info": {string(cd)},
	}
	resp, err := client.Post(COMMITURL, "application/x-www-form-urlencoded; charset=UTF-8", strings.NewReader(params.Encode()))
	if err != nil {
		log.Println("commit: client.Post: " + err.Error())
		return commitResult, err
	}
	body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &commitResult)
	if err != nil {
		log.Println("commit: json.Unmarshal: " + err.Error())
		return commitResult, err
	}
	return commitResult, nil
}

/**
利用默认的信息构造提交数据
*/
func genCommitData(reportData *model.ReportData) model.CommitData {
	modelData := &model.ModelData{}
	mapper.Mapper(reportData, modelData)
	commitData := model.CommitData{
		Model: *modelData,
	}
	return commitData
}
