package api

import (
	"encoding/json"
	"strconv"
	"util/http"
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
