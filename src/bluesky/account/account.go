package account

import (
	"util/bytes"
	"util/cryptor/md5"
)

type BotAccount struct {
	Uin      int
	Password string
}

func (a *BotAccount) Md5Password() []byte {
	return md5.StrToMd5Bytes(a.Password)
}

func (a *BotAccount) Md5UinPassword() []byte {
	builder := bytes.CreateBuilderByData(a.Md5Password())
	_ = builder.WriteLong(a.Uin)
	return md5.BsToMd5Bytes(builder.Bytes())
}
