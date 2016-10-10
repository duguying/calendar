package sms

import (
	"fmt"
	"github.com/gogather/com"
	"github.com/gogather/com/log"
	"github.com/gogather/yunpian"
	"path/filepath"
)

var message *yunpian.SMS

func getHome() string {
	home, err := com.Home()
	if err != nil {
		log.Fatalln("Can NOT find user path!")
	}
	return home
}

func getYunPianApiKey() string {
	home := getHome()
	path := filepath.Join(home, ".calendar", "api.key")
	key, err := com.ReadFileString(path)
	if err != nil {
		log.Redf("加载 API key 失败, 确认 %s 中存在正确的key\n", path)
		return ""
	} else {
		return com.Strim(key)
	}
}

func getSMSTpl() string {
	home := getHome()
	path := filepath.Join(home, ".calendar", "sms.tpl")
	tpl, err := com.ReadFileString(path)
	if err != nil {
		log.Redf("加载 短信模板 失败, 请确认 %s 存在\n", path)
		return ""
	} else {
		return tpl
	}
}

func init() {
	apikey := getYunPianApiKey()
	message = yunpian.NewSMS("https://sms.yunpian.com/v1/sms/send.json", apikey)
}

func SendSMS(phone string, time string, event string) {
	tpl := getSMSTpl()
	text := fmt.Sprintf(tpl, time, event)
	info, err := message.Send(phone, text)
	if err != nil {
		log.Redln(err)
	} else {
		log.Println(info)
	}
}
