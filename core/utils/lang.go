package utils

import (
	"encoding/json"
	"fmt"
	"net/url"
	"reflect"
	"sort"
	"strings"
)

func ObjectToMap(object interface{}) map[string]interface{} {
	t := reflect.TypeOf(object)
	v := reflect.ValueOf(object)
	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

func JsonToMap(jsonStr string) map[string]interface{} {
	var mapResult map[string]interface{}
	json.Unmarshal([]byte(jsonStr), &mapResult)
	return mapResult
}

func JsonToMapValues(jsonStr string) map[string]interface{} {
	var mapResult map[string]interface{}
	_ = json.Unmarshal([]byte(jsonStr), &mapResult)

	result := make(map[string]interface{}, len(mapResult))
	for k, v := range mapResult {
		if v == nil {
			continue
		}
		if isBasicType(v) {
			result[k] = v
		} else {
			bytes, _ := json.Marshal(v)
			result[k] = string(bytes)
		}
	}
	return result
}

func XmlToMapValues(ref interface{}) map[string]interface{} {
	mapResult := make(map[string]interface{})
	refTypes := reflect.TypeOf(ref)
	refValues := reflect.ValueOf(ref)
	if refTypes.Kind() == reflect.Ptr {
		refTypes = refTypes.Elem()
		refValues = refValues.Elem()
	}

	for i := 0; i < refTypes.NumField(); i++ {
		fieldType := refTypes.Field(i)
		fieldTag := fieldType.Tag
		field := refValues.Field(i)
		i2 := field.Interface()
		fmt.Print(reflect.ValueOf(&i2))

		if field.Kind() == reflect.Struct {
			values := XmlToMapValues(i2)
			for key, value := range values {
				mapResult[key] = value
			}
		}
		tagName := fieldTag.Get("xml")
		if tagName == "" || tagName == "-" || tagName == "xml" {
			continue
		}

		tagName = strings.Split(tagName, ",")[0]
		mapResult[tagName] = refValues.FieldByName(fieldType.Name).Interface()
	}
	return mapResult
}

func MapFilterNullOrEmptry(values map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{}, len(values))
	for k, v := range values {
		if v == nil || v == "" {
			continue
		}
		result[k] = v
	}
	return result
}

func KeysOrdered(keys []string, desc bool) []string {
	if desc {
		sort.Sort(sort.Reverse(sort.StringSlice(keys)))
	} else {
		sort.Strings(keys)
	}
	return keys
}

func Keys(values map[string]interface{}) []string {
	keys := make([]string, len(values))

	i := 0
	for k, _ := range values {
		keys[i] = k
		i++
	}
	return keys
}

func ConnectEncode(values map[string]interface{}, keys []string, divider string) string {
	return Connect(values, keys, divider, func(v interface{}) string {
		if isBasicType(v) {
			return fmt.Sprintf("%s", v)
		} else {
			bytes, _ := json.Marshal(v)
			return url.QueryEscape(fmt.Sprintf("%s", string(bytes)))
		}
	})
}

func Connect(values map[string]interface{}, keys []string, divider string, handleValue func(interface{}) string) string {
	result := ""
	for i, key := range keys {
		if i != 0 {
			result += divider
		}
		result += fmt.Sprintf(`%s=%s`, key, handleValue(values[key]))
	}
	return result
}

func ToJavaFieldName(str string) string {
	temp := strings.Split(str, "_")
	var upperStr string
	for y := 0; y < len(temp); y++ {
		vv := []rune(temp[y])
		if y != 0 {
			upperStr += FirstToUpper(string(vv))
		}
	}
	return temp[0] + upperStr
}

func ToGoFieldName(fname string) string {
	strs := strings.Split(fname, "_")
	result := ""
	for _, s := range strs {
		result += FirstToUpper(s)
	}
	return result
}

func FirstToUpper(str string) string {
	var upperStr string
	vv := []rune(str)
	for i := 0; i < len(vv); i++ {
		if i == 0 {
			vv[i] -= 32
			upperStr += string(vv[i])
		} else {
			upperStr += string(vv[i])
		}
	}
	return upperStr
}

func isBasicType(i interface{}) bool {
	switch i.(type) {
	case string:
		return true
	case int:
		return true
	case bool:
		return true
	case float32:
		return true
	}
	return false
}

// url参数转map
func QueryStrings(queryParams string) map[string]interface{} {
	result := make(map[string]interface{})
	values := strings.Split(queryParams, "&")
	for _, value := range values {
		kv := strings.Split(value, "=")
		if len(kv) != 2 {
			continue
		}
		key := kv[0]
		if unescape, err := url.QueryUnescape(kv[1]); err == nil {
			result[key] = unescape
		}
	}
	return result
}

// url参数转对象
func QueryParams(queryParams string, v interface{}) error {
	values := strings.Split(queryParams, "&")
	result := "{"
	size := len(values)

	for index, value := range values {
		kv := strings.Split(value, "=")
		if len(kv) != 2 {
			continue
		}
		key := kv[0]
		if unescape, err := url.QueryUnescape(kv[1]); err == nil {
			val := ""
			if (strings.HasPrefix(unescape, "[") && strings.HasSuffix(unescape, "]")) || (strings.HasPrefix(unescape, "(") && strings.HasSuffix(unescape, ")")) {
				val = unescape
			} else {
				val = fmt.Sprintf(`"%s"`, unescape)
			}
			spr := ""
			if index < (size - 1) {
				spr = ","
			}
			result += fmt.Sprintf(`"%s":%s%s`, key, val, spr)
		}
	}
	result += "}"
	return json.Unmarshal([]byte(result), v)
}
