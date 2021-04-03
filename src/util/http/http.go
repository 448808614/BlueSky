package http

import (
	"io/ioutil"
	"log"
	"net/http"
)

type Respond struct {
	Code int
	Body []byte
}

// Http组件
func Get(url string) *Respond {
	return GetHC(url, make(map[string]string), make(map[string]string))
}

func GetC(url string, cookie map[string]string) *Respond {
	return GetHC(url, make(map[string]string), cookie)
}

func GetH(url string, header map[string]string) *Respond {
	return GetHC(url, header, make(map[string]string))
}

func GetHC(url string, header map[string]string, cookie map[string]string) *Respond {
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Default().Println("[Request0x9]", err)
		return nil
	}
	for key, value := range header {
		request.Header.Add(key, value)
	}
	for key, value := range cookie {
		cookie := &http.Cookie{Name: key, Value: value}
		request.AddCookie(cookie)
	}
	resp, err := client.Do(request)
	if err != nil {
		log.Default().Println("[Request0xA]", err)
		return nil
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Default().Println("[Request0xB]", "The")
		return &Respond{resp.StatusCode, nil}
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Default().Println("[Request0xC]", err)
		return nil
	}
	return &Respond{resp.StatusCode, body}
}
