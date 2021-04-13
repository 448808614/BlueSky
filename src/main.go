package main

import "C"
import (
	"bluesky"
)

var BotMap = map[int]*bluesky.BlueSky{}

func main() {
	bot := bluesky.NewBot(1372362033, "911586abc", false)

	println(bot.Login())

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
