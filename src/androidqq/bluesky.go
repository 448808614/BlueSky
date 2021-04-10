package androidqq

import (
	"androidqq/account"
	"androidqq/env"
	"androidqq/record"
	"androidqq/tlv"
	"api"
	"time"
)

// BotType
const (
	// NoLogin 未登录
	NoLogin byte = iota
)

type BlueSky struct {
	Account *account.BotAccount
	// 机器人状态
	BotStatus    byte
	ProtocolInfo *api.ProtocolInfo
	Record       *record.BotRecord
	Tlv          *tlv.Tlv
}

// TODO("登录返回包需要TLV-16A")

//NewBot uin 账号, password 密码, isHd 是否Hd登录
func NewBot(uin int, password string, isHd bool) *BlueSky {
	botAccount := account.BotAccount{Uin: uin, Password: password}
	botRecord := record.BotRecord{
		InitTime: int(time.Now().Unix()),
		KeyMap:   map[int]*record.Key{},
	}
	android := env.Android{}
	protocolInfo := api.GetProtocolInfo(isHd)
	t := tlv.Tlv{
		BotAccount:   &botAccount,
		ProtocolInfo: protocolInfo,
		Record:       &botRecord,
		Android:      &android,
	}
	return &BlueSky{
		Account:      &botAccount,
		BotStatus:    NoLogin,
		Record:       &botRecord,
		ProtocolInfo: protocolInfo,
		Tlv:          &t,
	}
}
