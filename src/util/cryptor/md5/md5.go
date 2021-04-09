package md5

import (
	"crypto/md5"
)

func ToMd5BytesV2(s string) []byte {
	m := md5.New()
	m.Write([]byte(s))
	return m.Sum(nil)
}

func ToMd5Bytes(bs []byte) []byte {
	m := md5.New()
	m.Write(bs)
	return m.Sum(nil)
}
