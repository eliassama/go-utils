package convert

import (
	"encoding/json"
	"strings"
)

// Struct2Map 结构体转 map，注意：这个方法对于数字是直接转换为 float
//
//goland:noinspection GoUnusedExportedFunction
func Struct2Map(data interface{}, value any) error {
	dataByte, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(dataByte, &value); err != nil {
		return err
	}

	return err
}

// Struct2MapByNumberFirst 结构体转 map，注意：这个方法对于数字是先尝试转换为 int
//
//goland:noinspection GoUnusedExportedFunction
func Struct2MapByNumberFirst(data interface{}, value any) error {
	dataByte, err := json.Marshal(data)
	if err != nil {
		return err
	}

	decoder := json.NewDecoder(strings.NewReader(string(dataByte)))
	decoder.UseNumber()
	if err = decoder.Decode(&value); err != nil {
		return err
	}

	return err
}
