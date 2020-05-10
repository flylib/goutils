package encrypt

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

/*
 #加密过程
（1）A生成一对密钥（公钥和私钥），私钥不公开，A自己保留。公钥为公开的，任何人可以获取。
（2）A传递自己的公钥给B，B用A的公钥对消息进行加密。
（3）A接收到B加密的消息，利用A自己的私钥对消息进行解密。*/

//var publicKey string = `-----BEGIN PUBLIC KEY-----
//MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDPb8U4rEMMK8jk2/uGeYN7CLGi
//vXSKWUc2yIH8rUY05k4Sb9+lgvdhXuDr8UK4PYkyqbJ1UiT61Q/ZnF5D405oS2Jl
//mYtDep7TvHEsdd2xS3kX3KMu0O2TCGfKUT1YwIhPiqOQLNhqoJgf6KCSfjekx9nc
//xD0M5YzeCVUAwl59RQIDAQAB
//-----END PUBLIC KEY-----`

//var privateKey string = `-----BEGIN RSA PRIVATE KEY-----
//MIICWwIBAAKBgQDPb8U4rEMMK8jk2/uGeYN7CLGivXSKWUc2yIH8rUY05k4Sb9+l
//gvdhXuDr8UK4PYkyqbJ1UiT61Q/ZnF5D405oS2JlmYtDep7TvHEsdd2xS3kX3KMu
//0O2TCGfKUT1YwIhPiqOQLNhqoJgf6KCSfjekx9ncxD0M5YzeCVUAwl59RQIDAQAB
//AoGAGYYnPlHz7gt1LLPkvyc0hm8LbHrjXCKgIJ2LYQvxF5E/CgW5/yOeTNzf0Chf
//jUwFFbbLvqPc6QBOcvhKoQ/XFcVsqSMur06pyjZ78gV2vaRedbyU6QsGUQcYCqaz
//cCzKaWpz4pMuMKfJ+feYLOHfx0ULVU4tUQ5CGQp8OXlHSgECQQDw5e0Gsk54HyKz
//1qpybvYVHXdMkYAwn1PPSc8luZim33eWqBH4AQ+M66DngAWmCFG0rZakE8UiEDx4
//D15vaQsnAkEA3HDVh+64JD3izXzHDEwXz+DkyJ61vUdC3GJ/Rfztzqgjey3fgYPu
//slIHFN5gVcBN16ln2v+nZWshKwJ19XlnswJAefNb56zybnsMnVAJ335u00ekcj2i
//UHsH+YMa+7UWIzwzlTAmUI9w6N0MCsXTljbV7gqGnS9o95KSmhDltK7PtwJAfWYO
//gjoxNCSkPWKq1HsA3LdBTkLCfb7o8PdzETw1h2ascGkDCOklQvlYn+10fbNcVL9A
//nhrqfc34W0AWHCMI8wJAfirKhg9h2HdImJhWMkmKNWLqS/dlmGQI0/vlnl6qP2HE
//05mYXaziWxabzL9oYbtMu9z/4HmW6NmQUDguqp5KwA==
//-----END RSA PRIVATE KEY-----`

// 加密
func RsaEncrypt(origData []byte, publicKey string) ([]byte, error) {
	block, _ := pem.Decode([]byte(publicKey))
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

// 解密
func RsaDecrypt(ciphertext []byte, privateKey string) ([]byte, error) {
	block, _ := pem.Decode([]byte(privateKey))
	if block == nil {
		return nil, errors.New("private key error!")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}
