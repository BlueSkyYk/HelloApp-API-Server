package utils

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

//Bmob开发者账号
const (
	appId  = "340be399c1e866bc3a4045f846d6581e"
	apiKey = "721af9063f2ffac5ddb24b32c7684ca7"

	//登录短信模板名称
	loginTemplateName = "用户登录"

	//获取短信验证码地址
	getSmsCodeUrl = "https://api2.bmob.cn/1/requestSmsCode"
	//验证短信验证码地址
	verSmsCodeUrl = "https://api2.bmob.cn/1/verifySmsCode/"
)

//获取短信验证码
func GetSmsCode(phone string) error {
	return nil
	param := map[string]string{
		"mobilePhoneNumber": phone,
		"template":          loginTemplateName,
	}
	//格式化Json
	data, err := json.Marshal(&param)
	if err != nil {
		return errors.New("获取短信验证码失败")
	}
	client := &http.Client{}
	req, err := http.NewRequest("POST", getSmsCodeUrl, strings.NewReader(string(data)))
	if err != nil {
		return errors.New("获取短信验证码失败")
	}
	//设置请求头
	req.Header.Set("X-Bmob-Application-Id", "")
	req.Header.Set("X-Bmob-REST-API-Key", apiKey)
	req.Header.Set("Content-Type", "application/json")
	//发送请求
	res, err := client.Do(req)
	//关闭连接
	defer res.Body.Close()
	if err != nil {
		return errors.New("获取短信验证码失败")
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil || res.StatusCode != 200 {
		return errors.New("获取短信验证码失败")
	}
	resultMap := map[string]string{}
	json.Unmarshal(body, &resultMap)
	//ssid:=resultMap["ssid"]
	return nil
}

//验证短信验证码
func VerSmsCode(phone string, smsCode string) error {
	return nil
	param := map[string]string{
		"mobilePhoneNumber": phone,
	}
	//格式化Json
	data, err := json.Marshal(&param)
	if err != nil {
		return errors.New("短信验证失败")
	}
	client := &http.Client{}
	req, err := http.NewRequest("POST", verSmsCodeUrl+smsCode, strings.NewReader(string(data)))
	if err != nil {
		return errors.New("短信验证失败")
	}
	//设置请求头
	req.Header.Set("X-Bmob-Application-Id", appId)
	req.Header.Set("X-Bmob-REST-API-Key", apiKey)
	req.Header.Set("Content-Type", "application/json")
	//发送请求
	res, err := client.Do(req)
	//关闭连接
	defer res.Body.Close()
	if err != nil {
		return errors.New("短信验证失败")
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return errors.New("短信验证失败")
	}
	resultMap := map[string]string{}
	json.Unmarshal(body, &resultMap)
	if ("ok" != resultMap["msg"]) {
		msg := resultMap["error"]
		if msg == "" {
			msg = "短信验证失败"
		}
		return errors.New(msg)
	}
	return nil
}
