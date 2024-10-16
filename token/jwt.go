package tokenutils

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

//#jwt由以下三部分构成：
//* Header:头部 （对应：Header）
//* Claims:声明 (对应：Payload)
//* Signature:签名 (对应：Signature)

//#Claims 组成部分
//1. aud 接收该JWT的一方
//2. exp 过期时间.通常与Unix UTC时间做对比过期后token无效
//3. jti 是自定义的id号
//4. iat 签名发行时间.
//5. iss 是签名的发行者.
//6. nbf 这条token信息生效时间.这个值可以不设置,但是设定后,一定要大于当前Unix UTC,否则token将会延迟生效.
//7. sub 该JWT所面向的用户

type SigningMethod string

/*
SHA-256、SHA-384 和 SHA-512 是 SHA-2（Secure Hash Algorithm 2）的三种常见变体。它们都是密码学哈希算法，
用于生成一个固定长度的哈希值（也称为消息摘要）。它们的主要用途是数据完整性验证、数字签名和密码存储等。
*/
const (
	SHA256 SigningMethod = "SHA-256"
	SHA384 SigningMethod = "SHA-384"
	SHA512 SigningMethod = "SHA-512"
)

type JWT struct {
	signatureKey []byte
	signMethod   SigningMethod
}

func NewJWT(signingKey string, signMethod SigningMethod) *JWT {
	return &JWT{[]byte(signingKey), signMethod}
}

// 生成认证
func (j *JWT) CreateJWT(payload map[string]interface{}, expire time.Duration) (string, error) {
	if payload == nil || len(payload) == 0 {
		return "", errors.New("empty payload ") //不签发空内容
	}
	now := time.Now()
	payload["iat"] = now.Unix()             //签发时间
	payload["exp"] = now.Add(expire).Unix() //过期时间
	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(payload)).SignedString(j.signatureKey)
}

// 解析认证
func (j *JWT) ParseJWT(token string) (map[string]interface{}, error) {
	result, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return j.signatureKey, nil
	})
	if err != nil {
		return nil, err
	}
	return result.Claims.(jwt.MapClaims), err
}
