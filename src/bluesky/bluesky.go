package bluesky

import (
	"api"
	"bluesky/account"
	"bluesky/env"
	"bluesky/record"
	"bluesky/tlv"
	"container/list"
	"time"
	"util/tcp"
)

// BotType
const (
	// NoLogin 未登录
	NoLogin byte = iota
	// Fail 因为某种原因导致登录失败（看日志）
	Fail byte = iota
	// HasLogin 已登录,但没有上线
	HasLogin
	// Online 已上线
	Online
)

type BlueSky struct {
	Account *account.BotAccount
	// 机器人状态
	BotStatus    byte
	ProtocolInfo *api.ProtocolInfo
	Record       *record.BotRecord
	Tlv          *tlv.Tlv

	// 与腾讯服务器的连接器
	client *tcp.Tcp
	// 接解包器
	receiver *Receiver
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

func (r *BlueSky) AddWaiter(cmd string, seq int) *list.Element {
	pack := Packet{Cmd: cmd, Seq: seq, isOver: false}
	return r.receiver.waiter.PushFront(&pack)
}

func (r *BlueSky) WaitPacket(elem *list.Element) *Packet {
	waiter, ok := elem.Value.(*Packet)
	if ok {
		for waiter.isOver == false {
			if waiter.isOver == true {
				break
			}
		}
		return waiter
	} else {
		panic("这个屑玩意不是一个接包器~")
	}
}
