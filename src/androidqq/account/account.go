package account

import (
	"util/cryptor"
	"util/packet"
)

type BotAccount struct {
	Uin      int
	Password string
}

func (a *BotAccount) Md5Password() []byte {
	return cryptor.ToMd5BytesV2(a.Password)
}

func (a *BotAccount) Md5UinPassword() []byte {
	builder := packet.CreateBuilderByData(a.Md5Password())
	_ = builder.WriteLong(a.Uin)
	return cryptor.ToMd5Bytes(builder.Bytes())
}
