package http

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"util/packet"
)

type Http struct {
	cookie, header map[string]string
}

type Respond struct {
	Code int
	Body []byte
}

func CreateHttp() *Http {
	return &Http{
		header: make(map[string]string),
		cookie: make(map[string]string),
	}
}

// Http组件
func (http *Http) addHeader(key string, value string) {
	http.header[key] = value
}

func (http *Http) addCookie(key string, value string) {
	http.cookie[key] = value
}

func (http *Http) Get(url string) *Respond {
	return sendGet(url, http.header, http.cookie)
}

func (http *Http) Post(url string, data string) *Respond {
	return sendPost(url, data, http.header, http.cookie)
}

func (http *Http) PostBin(url string, data []byte) *Respond {
	return sendPostBin(url, data, http.header, http.cookie)
}

func (http *Http) PostJson(url string, data string) *Respond {
	return sendPostJson(url, packet.ToByteArray(data), http.header, http.cookie)
}

func sendPostJson(url string, data []byte, header map[string]string, cookie map[string]string) *Respond {
	reader := bytes.NewReader(data)
	client := &http.Client{}
	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		log.Default().Println("[Request0x5]", err)
		return nil
	}
	for key, value := range header {
		request.Header.Add(key, value)
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	for key, value := range cookie {
		cookie := &http.Cookie{Name: key, Value: value}
		request.AddCookie(cookie)
	}
	resp, err := client.Do(request)
	if err != nil {
		log.Default().Println("[Request0x6]", err)
		return nil
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Default().Println("[Request0x7]", "The")
		return &Respond{resp.StatusCode, nil}
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Default().Println("[Request0x8]", err)
		return nil
	}
	return &Respond{resp.StatusCode, body}
}

func sendPostBin(url string, data []byte, header map[string]string, cookie map[string]string) *Respond {
	reader := bytes.NewReader(data)
	client := &http.Client{}
	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		log.Default().Println("[Request0x5]", err)
		return nil
	}
	for key, value := range header {
		request.Header.Add(key, value)
	}
	request.Header.Set("Connection", "Keep-Alive")
	for key, value := range cookie {
		cookie := &http.Cookie{Name: key, Value: value}
		request.AddCookie(cookie)
	}
	resp, err := client.Do(request)
	if err != nil {
		log.Default().Println("[Request0x6]", err)
		return nil
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Default().Println("[Request0x7]", "The")
		return &Respond{resp.StatusCode, nil}
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Default().Println("[Request0x8]", err)
		return nil
	}
	return &Respond{resp.StatusCode, body}
}

func sendPost(url string, data string, header map[string]string, cookie map[string]string) *Respond {
	contentType := " application/x-www-form-urlencoded"
	reader := strings.NewReader(data)
	client := &http.Client{}
	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		log.Default().Println("[Request0x5]", err)
		return nil
	}
	for key, value := range header {
		request.Header.Add(key, value)
	}
	request.Header.Set("Content-Type", contentType)
	for key, value := range cookie {
		cookie := &http.Cookie{Name: key, Value: value}
		request.AddCookie(cookie)
	}
	resp, err := client.Do(request)
	if err != nil {
		log.Default().Println("[Request0x6]", err)
		return nil
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Default().Println("[Request0x7]", "The")
		return &Respond{resp.StatusCode, nil}
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Default().Println("[Request0x8]", err)
		return nil
	}
	return &Respond{resp.StatusCode, body}
}

func sendGet(url string, header map[string]string, cookie map[string]string) *Respond {
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
