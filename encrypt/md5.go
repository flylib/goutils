package encrypt

import (
	"bytes"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"
)

//生成32位md5字串[小写]
func MD5(data []byte) string {
	h := md5.New()
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}

func Hash(str string) (hash string) {
	timestamp := []byte(strconv.FormatInt(time.Now().Unix(), 10))
	headers := bytes.Join([][]byte{[]byte(str), timestamp}, []byte{})
	hash = fmt.Sprintf("", sha256.Sum256(headers))
	return
}
