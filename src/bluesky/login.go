package bluesky

import (
	"api"
	"util/tcp"
)

const (
	// Fail 因为某种原因导致登录失败（看日志）
	Fail byte = iota
	// HasLogin 已登录
	HasLogin
)

func (s *BlueSky) Login() byte {
	server := api.GetTencentServer()
	if s.client == nil {
		s.client = tcp.CreateTcp(server.Host, server.Port)
		s.InitReceive()

	} else {
		return HasLogin
	}
	return Fail
}
