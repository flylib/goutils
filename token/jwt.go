package utils

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

const (
	ClaimsKey = "PlayLoad"
)

var (
	jwtPrivateKey = []byte("000_%fj1@aj188￥2020") //加密私钥
)

//生成认证
func GenJWT(m map[string]interface{}, expire time.Duration) (string, error) {
	if m == nil {
		return "", errors.New("jwt:empty content ") //不签发空内容
	}
	now := time.Now()
	m["iat"] = now.Unix()             //签名时间
	m["exp"] = now.Add(expire).Unix() //过期时间
	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(m)).SignedString(jwtPrivateKey)
}

//解析认证
func ParseJWT(v string) (map[string]interface{}, error) {
	if len(v) < 1 {
		return map[string]interface{}{}, errors.New("Invalid token")
	}
	token, err := jwt.Parse(v, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return jwtPrivateKey, nil
	})
	return token.Claims.(jwt.MapClaims), err
}
