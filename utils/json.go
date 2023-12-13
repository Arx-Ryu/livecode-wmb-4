package utils

import "encoding/json"

func ToJsonString(data interface{}) string {
	var jsonData, err = json.Marshal(data)
	if err != nil {
		panic(err)
	}
	return string (jsonData) + "\n"
}

func FromJsonString(jsonString string, result interface{}) {
	var err = json.Unmarshal([]byte(jsonString), result)
	if err != nil {
		panic(err)
	}
}