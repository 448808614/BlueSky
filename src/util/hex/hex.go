package hex

import (
	"encoding/hex"
	"strings"
)

// Bytes2Str 字节转hex
func Bytes2Str(src []byte) string {
	dst := make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(dst, src)
	return string(dst)
}

// Str2Bytes hex转字节
func Str2Bytes(hexStr string) []byte {
	hexStr = strings.Replace(hexStr, "\n", "", -1)
	hexStr = strings.Replace(hexStr, "\r", "", -1)
	hexStr = strings.Replace(hexStr, "\t", "", -1)
	hexStr = strings.Replace(hexStr, " ", "", -1)
	src := []byte(hexStr)
	dst := make([]byte, hex.DecodedLen(len(src)))
	_, err := hex.Decode(dst, src)
	if err != nil {
		return []byte{}
	}
	return dst
}
