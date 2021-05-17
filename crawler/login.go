package crawler

import (
	"auto-report/model"
	"context"
	"errors"
	"github.com/devfeel/mapper"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"regexp"
	"strings"
)

var (
	LTURL     = "http://xgsm.hitsz.edu.cn/zhxy-xgzs/xg_mobile/shsj/common"
	LOGINURL  = "https://sso.hitsz.edu.cn:7002/cas/login?service=http://xgsm.hitsz.edu.cn/zhxy-xgzs/common/casLogin?params=L3hnX21vYmlsZS94c0hvbWU="
	POSTID    = "http://xgsm.hitsz.edu.cn/zhxy-xgzs/xg_mobile/xs/csh"
	REPORT    = "http://xgsm.hitsz.edu.cn/zhxy-xgzs/xg_mobile/xs/getYqxx"
	COMMITURL = "http://xgsm.hitsz.edu.cn/zhxy-xgzs/xg_mobile/xs/saveYqxx"

	JW_Mirror   = "https://185.245.2.16:7002"
	XGSM_Mirror = "http://185.245.2.16:7004"
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

// set mirror trickily
func setMirror(client http.Client) http.Client {
	client.Transport = &http.Transport{DialContext: func(ctx context.Context, network string, addr string) (net.Conn, error) {
		_, port, err := net.SplitHostPort(addr)
		if err != nil {
			return nil, err
		}
		var target string
		// tricky
		if port == "7002" {
			target = JW_Mirror
		} else {
			target = XGSM_Mirror
		}
		u, err := url.Parse(target)
		if err != nil {
			panic(err)
		}
		ip := u.Hostname()
		port = u.Port()
		if port == "" {
			if u.Scheme == "https" {
				port = "443"
			} else {
				port = "80"
			}
		}
		if net.ParseIP(ip) == nil {
			ips, err := net.LookupHost(ip)
			if err != nil {
				return nil, err
			}
			ip = ips[0]
		}
		return net.Dial(network, net.JoinHostPort(ip, port))
	}}
	return client
}

func login(account, password string) (http.Client, error) {
	jar, _ := cookiejar.New(nil)
	var client = http.Client{
		Jar: jar,
	}
	// github action
	if value := os.Getenv("CI"); value == "true" {
		client = setMirror(client)
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
		log.Println("登录失败！用户名或密码错误！")
		return client, errors.New("login error")
	}
	return client, nil
}
