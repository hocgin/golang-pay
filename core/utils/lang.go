package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"
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
		if tagName == "" || tagName == "xml" {
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

func Convert(value string, ref reflect.Value) (reflect.Value, error) {
	refType := ref.Type()
	switch refType {
	case reflect.TypeOf(0):
		val, err := strconv.Atoi(value)
		return reflect.ValueOf(val), err
	case reflect.TypeOf(int8(0)):
		val, err := strconv.ParseInt(value, 10, 8)
		return reflect.ValueOf(int8(val)), err
	case reflect.TypeOf(int16(0)):
		val, err := strconv.ParseInt(value, 10, 16)
		return reflect.ValueOf(int16(val)), err
	case reflect.TypeOf(int32(0)):
		val, err := strconv.ParseInt(value, 10, 32)
		return reflect.ValueOf(int32(val)), err
	case reflect.TypeOf(int64(0)):
		val, err := strconv.ParseInt(value, 10, 64)
		return reflect.ValueOf(val), err
	case reflect.TypeOf(uint(0)):
		val, err := strconv.ParseUint(value, 10, 0)
		return reflect.ValueOf(uint(val)), err
	case reflect.TypeOf(uint8(0)):
		val, err := strconv.ParseUint(value, 10, 8)
		return reflect.ValueOf(uint8(val)), err
	case reflect.TypeOf(uint16(0)):
		val, err := strconv.ParseUint(value, 10, 16)
		return reflect.ValueOf(uint16(val)), err
	case reflect.TypeOf(uint32(0)):
		val, err := strconv.ParseUint(value, 10, 32)
		return reflect.ValueOf(uint32(val)), err
	case reflect.TypeOf(uint64(0)):
		val, err := strconv.ParseUint(value, 10, 64)
		return reflect.ValueOf(val), err
	case reflect.TypeOf(float32(0)):
		val, err := strconv.ParseFloat(value, 32)
		return reflect.ValueOf(float32(val)), err
	case reflect.TypeOf(float64(0)):
		val, err := strconv.ParseFloat(value, 64)
		return reflect.ValueOf(val), err
	case reflect.TypeOf(false):
		val, err := strconv.ParseBool(value)
		if err == nil {
			return reflect.ValueOf(val), nil
		}
		val2, err := strconv.Atoi(value)
		if err == nil {
			return reflect.ValueOf(val2 != 0), nil
		}
		return reflect.Value{}, err
	case reflect.TypeOf(time.Now()):
		val, err := time.Parse("2006-01-02 15:04:05", value)
		return reflect.ValueOf(val), err
	default:
	}
	return reflect.ValueOf(nil), errors.New("转换失败")
}

func InjectValue(v interface{}, params map[string]string) {
	//elem := reflect.ValueOf(v).Elem()
	ref := reflect.ValueOf(v).Elem()
	for i := 0; i < ref.NumField(); i++ {
		fieldType := ref.Type().Field(i)
		fieldTag := fieldType.Tag
		field := ref.Field(i)

		if field.Kind() == reflect.Struct {
			i2 := reflect.ValueOf(v)
			reflect.ValueOf(&i2).Elem()
			//InjectValue(&i2, params)
		}

		tagName := fieldTag.Get("json")
		if tagName == "" {
			tagName = strings.ToLower(fieldType.Name)
		}

		tagName = strings.Split(tagName, ",")[0]
		if value, ok := params[tagName]; ok {
			field := ref.FieldByName(fieldType.Name)
			if reflect.ValueOf(value).Type() == field.Type() {
				field.Set(reflect.ValueOf(value))
			} else if value, err := Convert(value, field); err == nil {
				field.Set(value)
			}
		}
	}
}
