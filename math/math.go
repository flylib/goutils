package math

//10 的n次方
func TenCube(n int) int {
	if n == 0 {
		return 1
	}
	return 10 * TenCube(n-1)
}
