package main

import "C"
import (
	"util/cryptor"
	"util/hex"
)

func main() {
	println(hex.Bytes2Str(cryptor.ToMd5Bytes([]byte("你好"))))
}

//export test
func test() string {
	return "Hello Go!!!"
}
