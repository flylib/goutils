package math

import (
	"math/rand"
)

//10 的n次方
func TenCube(n int) int {
	if n == 0 {
		return 1
	}
	return 10 * TenCube(n-1)
}

//@rts 概率分配 比如[0.1,0.3,0.5,0.1]
//@seed 随机种子
//按概率分配随机抽奖
func Lottery(rts []float64, seed int64) int {
	rand.Seed(seed)
	f := rand.Float64()
	var rangeMax float64
	for i := 0; i < len(rts); i++ {
		rangeMax += rts[i]

		if f <= rangeMax {
			return i + 1
		}
	}
	return 0
}
