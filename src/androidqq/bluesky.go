package androidqq

import (
	"androidqq/account"
	"androidqq/record"
	"androidqq/tlv"
	"api"
	"time"
)

// BotType
const (
	// 未登录
	NoLogin int = iota
)

type BlueSky struct {
	Account *account.BotAccount
	// 机器人状态
	BotStatus    int
	ProtocolInfo *api.ProtocolInfo
	Record       *record.BotRecord
	Tlv          *tlv.Tlv
}

//NewBot uin 账号, password 密码, isHd 是否Hd登录
func NewBot(uin int, password string, isHd bool) *BlueSky {
	botAccount := account.BotAccount{Uin: uin, Password: password}
	botRecord := record.BotRecord{
		InitTime: int(time.Now().Unix()),
	}

	return &BlueSky{
		Account:   &botAccount,
		BotStatus: NoLogin,
		Record:    &botRecord,
	}
}
