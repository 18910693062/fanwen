package util

import (
	"encoding/json"
	"strconv"
)

func StrtoInt(a string) int {
	ret, _ := strconv.Atoi(a)
	return ret
}

func InttoStr(a int) string {
	ret := strconv.Itoa(a)
	return ret
}

func jsontomap(a []byte) map[string]interface{} {
	var info = make(map[string]interface{})
	err := json.Unmarshal(a, &info)
	if err != nil {
		return nil
	}
	return info
}

func maptojson(a any) string {
	ret, err := json.Marshal(a)
	if err != nil {
		return "err"
	}
	return string(ret)
}
