package core

import (
	"bytes"
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
)

func Md5(data string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(data)))
}

// RSA 签名: use PKCS8 privateKey
func RsaSign(signContent string, privateKey string, hash crypto.Hash) string {
	shaNew := hash.New()
	shaNew.Write([]byte(signContent))
	hashed := shaNew.Sum(nil)
	priKey, err := GetPrivateKey(privateKey)
	if err != nil {
		panic(err)
	}
	signature, err := rsa.SignPKCS1v15(rand.Reader, priKey, hash, hashed)
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(signature)
}

// RSA 签名验证:
func RsaVerify(signContent string, publicKey string, verifySign string, hash crypto.Hash) bool {
	pubKey, err := GetPublicKey(publicKey)
	if err != nil {
		_ = fmt.Errorf("{} 签名公钥错误", publicKey)
		return false
	}
	signData, err := base64.StdEncoding.DecodeString(verifySign)
	if err != nil {
		_ = fmt.Errorf("{} 签名字符串解析错误", publicKey)
		return false
	}

	h := hash.New()
	h.Write([]byte(signContent))
	digest := h.Sum(nil)

	err = rsa.VerifyPKCS1v15(pubKey, hash, digest, signData)
	if err == nil {
		return true
	}
	_ = fmt.Errorf("签名验证失败")
	return false
}

func GetPublicKey(publicKey string) (*rsa.PublicKey, error) {
	publicKey = FormatPublicKey(publicKey)
	block, _ := pem.Decode([]byte(publicKey))
	if block == nil {
		return nil, errors.New("get public key error")
	}
	// x509 parse public key
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return pub.(*rsa.PublicKey), err
}

func GetPrivateKey(privateKey string) (*rsa.PrivateKey, error) {
	privateKey = FormatPrivateKey(privateKey)
	keyBytes := []byte(privateKey)
	block, _ := pem.Decode(keyBytes)
	if block == nil {
		panic(errors.New("private key error"))
	}

	// PKCS1
	pri, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err == nil {
		return pri, nil
	}

	// PKCS8
	priKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return priKey.(*rsa.PrivateKey), nil
}

// 格式化私钥
func FormatPrivateKey(privateKey string) string {
	buffer := new(bytes.Buffer)
	buffer.WriteString("-----BEGIN RSA PRIVATE KEY-----\n")
	rawLen, keyLen := 64, len(privateKey)
	raws, temp := keyLen/rawLen, keyLen%rawLen
	if temp > 0 {
		raws++
	}
	start, end := 0, 0+rawLen
	for i := 0; i < raws; i++ {
		if i == raws-1 {
			buffer.WriteString(privateKey[start:])
		} else {
			buffer.WriteString(privateKey[start:end])
		}
		buffer.WriteString("\n")
		start += rawLen
		end = start + rawLen
	}
	buffer.WriteString("-----END RSA PRIVATE KEY-----\n")
	return buffer.String()
}

// 格式化公钥
func FormatPublicKey(publicKey string) string {
	buffer := new(bytes.Buffer)
	buffer.WriteString("-----BEGIN PUBLIC KEY-----\n")
	rawLen, keyLen := 64, len(publicKey)
	raws, temp := keyLen/rawLen, keyLen%rawLen
	if temp > 0 {
		raws++
	}
	start, end := 0, 0+rawLen
	for i := 0; i < raws; i++ {
		if i == raws-1 {
			buffer.WriteString(publicKey[start:])
		} else {
			buffer.WriteString(publicKey[start:end])
		}
		buffer.WriteString("\n")
		start += rawLen
		end = start + rawLen
	}
	buffer.WriteString("-----END PUBLIC KEY-----\n")
	return buffer.String()
}
