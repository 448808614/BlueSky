package tlv

import (
	"androidqq/account"
	"androidqq/env"
	"androidqq/record"
	"api"
)

type Tlv struct {
	BotAccount   *account.BotAccount
	ProtocolInfo *api.ProtocolInfo
	Record       *record.BotRecord
	Android      *env.Android
}

var domains = []string{
	"tenpay.com",
	"qzone.qq.com",
	"qun.qq.com",
	"mail.qq.com",
	"openmobile.qq.com",
	"connect.qq.com",
	"qqweb.qq.com",
	"office.qq.com",
	"ti.qq.com",
	"mma.qq.com",
	"docs.qq.com",
	"gamecenter.qq.com",
	"qzone.com",
	"game.qq.com",
	"vip.qq.com",
	"weiyun.com",
	"lol.qq.com",
	"b.qq.com",
}
