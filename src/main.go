package main

import "C"
import (
	"androidqq"
	"api"
	"github.com/gogo/protobuf/proto"
	login "protocol/protobuf"
	"util/hex"
)

var BotMap = map[int]*androidqq.BlueSky{}

func main() {
	bot := androidqq.NewBot(1372362033, "911586abc", false)

	println(bot)

	api.GetProtocolInfo(false)

	abcd := login.DeviceReport{}

	abcd.BootId = []byte("你好")

	bytes, _ := proto.Marshal(&abcd)

	println(hex.Bytes2Str(bytes))
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
