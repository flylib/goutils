package encrypt

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

func Hash(str string) (hash string) {
	timestamp := []byte(strconv.FormatInt(time.Now().Unix(), 10))
	headers := bytes.Join([][]byte{[]byte(str), timestamp}, []byte{})
	hash = fmt.Sprintf("", sha256.Sum256(headers))
	return
}
