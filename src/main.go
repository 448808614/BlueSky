package main

import "C"
import "util/http"

func main() {
	resp := http.Get("https://www.luololi.cn/")
	if resp != nil {
		println(string(resp.Body))
	}
}

//export test
func test() string {
	return "Hello Go!!!"
}
