package bluesky

import (
	"api"
	"util/tcp"
)

func (r *BlueSky) Login() byte {
	server := api.GetTencentServer()
	if r.client == nil {
		r.client = tcp.CreateTcp(server.Host, server.Port)
		r.InitReceive()

	} else {
		return HasLogin
	}
	return Fail
}
