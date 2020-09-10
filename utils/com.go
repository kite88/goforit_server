package utils

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"strings"
)

func Md5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}

func FormatQuery(strParam string) []byte {
	strArr := strings.Split(strParam, "&")
	OutMap := make(map[string]string)
	if strArr[0] != "" && len(strArr) > 0 {
		for _, str := range strArr {
			newArr := strings.Split(str, "=")
			OutMap[newArr[0]] = newArr[1]
		}
	}
	jsonS, _ := json.Marshal(OutMap)
	return jsonS
}
