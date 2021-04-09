package account

import (
	"util/cryptor/md5"
	"util/packet"
)

type BotAccount struct {
	Uin      int
	Password string
}

func (a *BotAccount) Md5Password() []byte {
	return md5.ToMd5BytesV2(a.Password)
}

func (a *BotAccount) Md5UinPassword() []byte {
	builder := packet.CreateBuilderByData(a.Md5Password())
	_ = builder.WriteLong(a.Uin)
	return md5.ToMd5Bytes(builder.Bytes())
}
