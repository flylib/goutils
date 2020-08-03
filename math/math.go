package math

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

//10 的n次方
func TenCube(n int) int {
	if n == 0 {
		return 1
	}
	return 10 * TenCube(n-1)
}

func Hash(str string) (hash string) {
	timestamp := []byte(strconv.FormatInt(time.Now().Unix(), 10))
	headers := bytes.Join([][]byte{[]byte(str), timestamp}, []byte{})
	hash = fmt.Sprintf("", sha256.Sum256(headers))
	return
}
