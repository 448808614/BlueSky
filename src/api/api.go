package api

import (
	"encoding/json"
	"log"
	"strconv"
	"util/http"
	json2 "util/json"
)

// 云端API接口服务·类

const centerServer = "https://www.luololi.cn/?s="

type tencentServerRet struct {
	Ret  int    `json:"ret" gorm:"column:ret"`
	Msg  string `json:"msg" gorm:"column:msg"`
	Data struct {
		Ip []string `json:"ip" gorm:"column:ip"`
	} `json:"data" gorm:"column:data"`
}

type TencentServer struct {
	Host string
	Port int
}

type ProtocolInfo struct {
	MsfAppId  int
	SubAppId  int
	IpVersion int
	LocalId   int
}

// 获取腾讯QQ服务器
func GetTencentServer() *TencentServer {
	var host = "msfwifi.3g.qq.com"
	var port = 8080
	request := http.CreateHttp()
	resp := request.Get(centerServer + "Main.tencentServer")
	if resp != nil {
		ret := tencentServerRet{}
		err := json.Unmarshal(resp.Body, &ret)
		if err == nil {
			port, err = strconv.Atoi(ret.Data.Ip[1])
			if err == nil {
				host = ret.Data.Ip[0]
			} else {
				port = 8080
			}
		}
	}
	return &TencentServer{host, port}
}

func GetProtocolInfo(isHd bool) *ProtocolInfo {
	request := http.CreateHttp()
	resp := request.Get(centerServer + "Main.protocolInfo")
	if resp != nil {
		var latestKey = "latest"
		if isHd {
			latestKey = "hdLatest"
		}
		body := resp.Body
		data := json2.Get(body, "data")
		version := json2.GetString(data, latestKey)
		info := json2.Get(data, version)
		return &ProtocolInfo{
			MsfAppId:  json2.GetInt(info, "appid"),
			SubAppId:  json2.GetInt(info, "subAppid"),
			IpVersion: json2.GetInt(info, "ipVersion"),
			LocalId:   json2.GetInt(info, "lId"),
		}
	}
	log.Default().Println("登录失败，无法获取协议信息")
	return nil
}
