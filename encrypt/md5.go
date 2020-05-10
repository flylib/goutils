package encrypt

import (
	"crypto/md5"
	"encoding/hex"
)

//生成32位md5字串[小写]
func MD5(data []byte) string {
	h := md5.New()
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}
