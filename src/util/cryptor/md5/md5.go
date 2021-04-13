package md5

import (
	"crypto/md5"
)

func ToMd5Bytes(data interface{}) []byte {
	switch v := data.(type) {
	case *string:
		return StrToMd5Bytes(*v)
	case string:
		return StrToMd5Bytes(v)
	case *[]byte:
		return BsToMd5Bytes(*v)
	case []byte:
		return BsToMd5Bytes(v)
	}
	panic("未知的玩意，转不了MD5的嗷~~")
}

func StrToMd5Bytes(s string) []byte {
	m := md5.New()
	m.Write([]byte(s))
	return m.Sum(nil)
}

func BsToMd5Bytes(bs []byte) []byte {
	m := md5.New()
	m.Write(bs)
	return m.Sum(nil)
}
