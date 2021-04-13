package api

import (
	"log"
	"util/http"
	"util/json"
)

type ProtocolInfo struct {
	MsfAppId        int
	SubAppId        int
	IpVersion       int
	LocalId         int
	PingVersion     int
	SSoVersion      int
	DbVersion       int
	MsfSSoVersion   int
	MiscBitmap      int
	SubSigMap       int
	TgtVersion      int
	PackageName     string
	PackageVersion  string
	MsfSdkMd5       string
	BuildTime       int
	BuildVersion    string
	ProtocolVersion string
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
		data := json.Get(body, "data")
		version := json.GetString(data, latestKey)
		info := json.Get(data, version)
		return &ProtocolInfo{
			MsfAppId:        json.GetInt(info, "appid"),
			SubAppId:        json.GetInt(info, "subAppid"),
			IpVersion:       json.GetInt(info, "ipVersion"),
			LocalId:         json.GetInt(info, "lId"),
			PingVersion:     json.GetInt(info, "pingVersion"),
			SSoVersion:      json.GetInt(info, "ssoVer"),
			DbVersion:       json.GetInt(info, "dbVer"),
			MsfSSoVersion:   json.GetInt(info, "msfSsoVer"),
			MiscBitmap:      json.GetInt(info, "miscBitmap"),
			SubSigMap:       json.GetInt(info, "subSigMap"),
			TgtVersion:      json.GetInt(info, "tgtgVer"),
			PackageName:     json.GetString(info, "packageName"),
			PackageVersion:  json.GetString(info, "packageVersion"),
			MsfSdkMd5:       json.GetString(info, "sdkMd5"),
			BuildTime:       json.GetInt(info, "buildTime"),
			BuildVersion:    json.GetString(info, "buildVersion"),
			ProtocolVersion: json.GetString(info, "agreementVersion"),
		}
	}
	log.Default().Println("登录失败，无法获取协议信息")
	return nil
}
