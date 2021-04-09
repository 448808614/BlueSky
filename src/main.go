package main

import "C"
import (
	"androidqq"
	"api"
)

var BotMap = map[int]*androidqq.BlueSky{}

func main() {
	bot := androidqq.NewBot(1372362033, "911586abc", false)

	println(bot)

	api.GetProtocolInfo(false)

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
