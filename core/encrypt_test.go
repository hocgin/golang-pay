package core

import (
	"testing"
)

func TestMd5(t *testing.T) {
	md5 := MD5("hocgin")
	if len(md5) != 32 {
		panic("len(md5) != 32")
	}
}
