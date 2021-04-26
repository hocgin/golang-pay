package sign

import (
	"github.com/hocgin/golang-pay/core"
	"strings"
)

type MD5 struct {
}

func (_ *MD5) Sign(data string, privateKey string) string {
	return core.MD5(data)
}

func (_ *MD5) Verify(data string, publicKey string, sign string) bool {
	return strings.ToUpper(core.MD5(data)) == strings.ToUpper(sign)
}
func (_ *MD5) Name() string {
	return "MD5"
}
