package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"

	"github.com/go-ini/ini"
)

type dnspodMsg struct {
	Status status `json:"status"`
}

type status struct {
	Code string `json:"code"`
	Msg  string `json:"message"`
}

type config struct {
	Dndpod dnspod `ini:"Dnspod"`
	Email  email  `ini:"Email"`
}
type dnspod struct {
	Token     string `ini:"token"`
	Format    string `ini:"format"`
	Domainid  string `ini:"domainid"`
	Recordid  string `ini:"recordid"`
	Subdomain string `ini:"subdomain"`
}
type email struct {
}

const (
	ipfile = "/tmp/ip"
	ipurl  = "http://ifconfig.me"
)

func init() {
	_, err := os.Stat(ipfile)
	if os.IsNotExist(err) {
		file, _ := os.Create(ipfile)
		defer file.Close()
	}
}

func main() {
	oldip, _ := ioutil.ReadFile(ipfile)
	ip := getip()
	if ip == string(oldip) {
		os.Exit(0)
	}
	config := getconfig()
	ddns(config.Dndpod)
	os.Truncate(ipfile, 0)
	ioutil.WriteFile(ipfile, []byte(ip), 0644)
	// fmt.Println(ip)
}

// 获取当前公网ip
func getip() string {
	c := &http.Client{}
	request, err := http.NewRequest("GET", ipurl, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	response, _ := c.Do(request)
	ip, _ := ioutil.ReadAll(response.Body)
	return string(ip)
}

// 修改dnspod域名
func ddns(dnspodinfo dnspod) *status {
	v := url.Values{}
	c := &http.Client{}
	v.Set("login_token", dnspodinfo.Token)
	v.Add("format", dnspodinfo.Format)
	v.Add("domain_id", dnspodinfo.Domainid)
	v.Add("record_id", dnspodinfo.Recordid)
	v.Add("record_line", "默认")
	v.Add("sub_domain", dnspodinfo.Subdomain)
	body := ioutil.NopCloser(strings.NewReader(v.Encode()))
	request, err := http.NewRequest("POST", "https://dnsapi.cn/Record.Ddns", body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	request.Header.Add("Content-type", "application/x-www-form-urlencoded")
	request.Header.Add("Accept", "text/json")
	response, _ := c.Do(request)
	msg := dnspodMsg{}
	result, _ := ioutil.ReadAll(response.Body)
	// result, _ = ioutil.ReadAll(request.Body)
	json.Unmarshal(result, &msg)
	fmt.Println(string(result))
	return &msg.Status
}

// 从配置文件读取dnspod信息和邮箱信息
func getconfig() *config {
	dnspodconfig := new(config)
	ini.MapTo(dnspodconfig, configpath())
	return dnspodconfig
}

//
func configpath() string {
	dir, _ := path.Split(os.Args[0])
	return dir + "config.ini"
}

// 有时间完善如下内容
// 修改失败 发送告警邮件到指定邮箱
func sendmail() {

}
