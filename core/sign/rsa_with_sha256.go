package sign

import (
	"crypto"
	"github.com/hocgin/golang-pay/core"
)

type RSAWithSHA256 struct {
}

func (_ *RSAWithSHA256) Name() string {
	return "RSAWithSHA256"
}

func (_ *RSAWithSHA256) Sign(data string, privateKey string) string {
	return core.RsaSign(data, privateKey, crypto.SHA256)
}

func (_ *RSAWithSHA256) Verify(data string, publicKey string, verifySign string) bool {
	return core.RsaVerify(data, publicKey, verifySign, crypto.SHA256)
}
