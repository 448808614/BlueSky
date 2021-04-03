package main

import "C"
import "api"

func main() {
	api.GetTencentServer()

}

//export test
func test() string {
	return "Hello Go!!!"
}
