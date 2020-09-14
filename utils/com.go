package utils

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/sony/sonyflake"
	"strconv"
	"strings"
	"github.com/rs/xid"
)

func Md5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}

func FormatQuery(strParam string, transitionField ...string) []byte {
	sets := make(map[string]struct{})
	for _, val := range transitionField {
		sets[val] = struct{}{}
	}
	strArr := strings.Split(strParam, "&")
	OutMap := make(map[string]interface{})
	if strArr[0] != "" && len(strArr) > 0 {
		for _, str := range strArr {
			newArr := strings.Split(str, "=")
			key := newArr[0]
			value := newArr[1]
			r, n := IsNumber(value)
			_, ok := sets[key]
			if r == true && ok == true {
				OutMap[key] = n
			} else {
				OutMap[key] = value
			}
		}
	}
	jsonS, _ := json.Marshal(OutMap)
	return jsonS
}

func IsNumber(s string) (bool, float64) {
	result, err := strconv.ParseFloat(s, 64)
	return err == nil, result
}

func GenMachineID() string {
	flake := sonyflake.NewSonyflake(sonyflake.Settings{})
	id, _ := flake.NextID()
	return fmt.Sprintf("%d", id)
}

func GenXid() string {
	id := xid.New()
	return id.String()
}
