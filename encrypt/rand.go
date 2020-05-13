package encrypt

import (
	"math/rand"
	"time"
)

const (
	RAND_NUM   = 0 // 纯数字
	RAND_LOWER = 1 // 小写字母
	RAND_UPPER = 2 // 大写字母
	RAND_ALL   = 3 // 数字、大小写字母
)

//生成随机码
func GenRandCode(inviteCodeLen int, targetKind int) string {
	ikind, kinds, result := targetKind, [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}, make([]byte, inviteCodeLen)
	isAll := targetKind > 2 || targetKind < 0
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < inviteCodeLen; i++ {
		if isAll {
			ikind = rand.Intn(3)
		}
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return string(result)
}
