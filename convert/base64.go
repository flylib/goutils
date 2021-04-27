package convert

import (
	"encoding/base64"
)

//base64编码
func Base64Encode(srcStr string) string {
	srcBytes := StringToBytes(srcStr)
	encoding := base64.StdEncoding.EncodeToString(srcBytes)
	return encoding
}

//base64解码
func Base64Decode(baseStr string) (string, error) {
	decodeBytes, err := base64.StdEncoding.DecodeString(baseStr)
	if err != nil {
		return "", err
	}
	return BytesToString(decodeBytes), nil
}
