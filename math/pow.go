package math

//浮点型n次方
func Pow(x float64, n int) float64 {
	if x == 0 {
		return 0
	}
	result := calPow(x, n)
	if n < 0 {
		result = 1 / result
	}
	return result
}

func calPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}

	// 向右移动一位
	result := calPow(x, n>>1)
	result *= result

	// 如果n是奇数
	if n&1 == 1 {
		result *= x
	}

	return result

}

//10 的n次方
func TenCube(n int) int {
	if n == 0 {
		return 1
	}
	return 10 * TenCube(n-1)
}
