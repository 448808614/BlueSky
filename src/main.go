package main

import "C"
import "util/http"

func main() {
	createHttp := http.CreateHttp()

	resp := createHttp.Get("https://blog.csdn.net/fyxichen/article/details/51258351")

	if resp != nil {
		println(string(resp.Body))
	}

}

//export test
func test() string {
	return "Hello Go!!!"
}
