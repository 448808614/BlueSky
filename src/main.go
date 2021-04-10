package main

import "C"
import (
	"androidqq"
	"github.com/gogo/protobuf/proto"
	login "protocol/protobuf"
	"util/hex"
)

var BotMap = map[int]*androidqq.BlueSky{}

func main() {
	bot := androidqq.NewBot(1372362033, "911586abc", false)

	println(bot)

	deviceReport := login.DeviceReport{
		Bootloader: []byte("unknown"),
		Version:    []byte("Linux version 4.19.113-perf-gb3dd08fa2aaa (builder@c5-miui-ota-bd143.bj) (clang version 8.0.12 for Android NDK) #1 SMP PREEMPT Thu Feb 4 04:37:10 CST 2021;"),
	}
	bytes, _ := proto.Marshal(&deviceReport)

	println(hex.Bytes2Str(bytes))

	println(hex.Bytes2Str(bot.Account.Md5UinPassword()))

}

//export love
func love() string {
	return "这个世界上的美丽多半大同小异，就好比我觉得好看的人，都像你。"
}

//export botSize
func botSize() int {
	// 获取机器人数量
	return len(BotMap)
}
