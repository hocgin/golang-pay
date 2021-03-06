package utils

import (
	"fmt"
	"testing"
	"time"
)

func TestSortMap(t *testing.T) {
	maps := make(map[string]interface{})
	maps["a"] = "a"
	maps["s"] = "s"
	maps["1"] = "1"

}
func TestToGoFieldName(t *testing.T) {
	fieldName := ToGoFieldName("gmt_create")
	fmt.Print(fieldName)
}
func TestToJavaFieldName(t *testing.T) {
	fieldName := ToJavaFieldName("gmt_create")
	fmt.Print(fieldName)
}
func TestFirstToUpper(t *testing.T) {
	fieldName := FirstToUpper("gmt")
	fmt.Print(fieldName)
}

type TestObject struct {
	Charset   string    `json:"charset"`
	Sign      string    `json:"sign"`
	Id        int32     `json:"id"`
	Enabled   bool      `json:"enabled"`
	CreatedAt time.Time `json:"created_at"`
}

//const (
//    timeFormat = "2006-01-02 15:04:05"
//)
//type Time time.Time
//func (t *Time) UnmarshalJSON(data []byte) (err error) {
//    now, err := time.ParseInLocation(`"`+timeFormat+`"`, string(data), time.Local)
//    *t = Time(now)
//    return
//}
