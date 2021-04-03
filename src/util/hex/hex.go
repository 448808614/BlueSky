package hex

import (
	"encoding/hex"
)

// 字节转hex
func Bytes2Str(src []byte) string {
	dst := make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(dst, src)
	return string(dst)
}

// hex转字节
func Str2Bytes(hexStr string) []byte {
	src := []byte(hexStr)
	dst := make([]byte, hex.DecodedLen(len(src)))
	_, err := hex.Decode(dst, src)
	if err != nil {
		return []byte{}
	}
	return dst
}
