package sort

import (
	"math/rand"
	"time"
)

//快速排序-沈旭
func QuickSortASC(li []int, left, right int) {
	if left >= right || right >= len(li) {
		return
	}
	i := left
	j := right
	rand.Seed(time.Now().Unix())
	r := rand.Intn(right-left) + left
	li[i], li[r] = li[r], li[i]
	tmp := li[i]
	for i < j {
		for i < j && li[j] >= tmp {
			j--
		}
		li[i] = li[j]
		for i < j && li[i] <= tmp {
			i++
		}
		li[j] = li[i]
	}
	li[i] = tmp
	QuickSortASC(li, left, i-1)
	QuickSortASC(li, i+1, right)
}

//快速排序-降序
func QuickSortDESC(li []int, left, right int) {
	if left >= right || right >= len(li) {
		return
	}
	i := left
	j := right
	rand.Seed(time.Now().Unix())
	r := rand.Intn(right-left) + left
	li[i], li[r] = li[r], li[i]
	tmp := li[i]
	for i < j {
		for i < j && li[j] <= tmp {
			j--
		}
		li[i] = li[j]
		for i < j && li[i] >= tmp {
			i++
		}
		li[j] = li[i]
	}
	li[i] = tmp
	QuickSortDESC(li, left, i-1)
	QuickSortDESC(li, i+1, right)
	return
}
