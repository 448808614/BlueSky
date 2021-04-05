package json

import "github.com/buger/jsonparser"

func Get(data []byte, key string) []byte {
	data, _, _, _ = jsonparser.Get(data, key)
	return data
}

func GetInt(data []byte, key string) int {
	i, _ := jsonparser.GetInt(data, key)
	return int(i)
}

func GetString(data []byte, key string) string {
	s, _ := jsonparser.GetString(data, key)
	return s
}
