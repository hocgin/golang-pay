package sign

import (
	"crypto"
	"crypto/hmac"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"strings"
)

const (
	// 私钥 PEMBEGIN 开头
	PEMBEGIN = "-----BEGIN RSA PRIVATE KEY-----\n"
	// 私钥 PEMEND 结尾
	PEMEND = "\n-----END RSA PRIVATE KEY-----"
	// 公钥 PEMBEGIN 开头
	PUBPEMBEGIN = "-----BEGIN PUBLIC KEY-----\n"
	// 公钥 PEMEND 结尾
	PUBPEMEND = "\n-----END PUBLIC KEY-----"
)

type RSA2 struct {
}

func (_ *RSA2) Sign(data string, privateKey string) string {
	hash := hmac.New(sha256.New, []byte(privateKey))

	if _, err := hash.Write([]byte(data)); err != nil {
		return ``
	}
	return strings.ToUpper(hex.EncodeToString(hash.Sum(nil)))
}

func (_ *RSA2) Verify(data string, publicKey string, sign string) bool {
	block, _ := pem.Decode([]byte(publicKey))
	if block == nil {
		panic(errors.New("public key error"))
	}
	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}

	hashed := sha256.Sum256([]byte(data))
	err = rsa.VerifyPKCS1v15(pubKey.(*rsa.PublicKey), crypto.SHA256, hashed[:], []byte(sign))
	if err != nil {
		panic(err)
	}
	return true
}

func (_ *RSA2) Name() string {
	return "RSA2"
}

func ParsePrivateKey(privateKey string) (*rsa.PrivateKey, error) {
	privateKey = FormatPrivateKey(privateKey)
	block, _ := pem.Decode([]byte(privateKey))
	if block == nil {
		return nil, errors.New("私钥信息错误！")
	}
	priKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return priKey, nil
}
func FormatPrivateKey(privateKey string) string {
	if !strings.HasPrefix(privateKey, PEMBEGIN) {
		privateKey = PEMBEGIN + privateKey
	}
	if !strings.HasSuffix(privateKey, PEMEND) {
		privateKey = privateKey + PEMEND
	}
	return privateKey
}
