package sign

import (
	"crypto"
	"github.com/hocgin/golang-pay/core"
)

type RSAWithSHA128 struct {
}

func (_ *RSAWithSHA128) Name() string {
	return "RSAWithSHA128"
}

func (_ *RSAWithSHA128) Sign(data string, privateKey string) string {
	return core.RsaSign(data, privateKey, crypto.SHA1)
}

func (_ *RSAWithSHA128) Verify(data string, publicKey string, verifySign string) bool {
	return core.RsaVerify(data, publicKey, verifySign, crypto.SHA1)
}
