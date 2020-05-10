package sort

import (
	"math/rand"
	"time"
)

//快速排序
func QuickSort(li []int, left, right int) {
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
	QuickSort(li, left, i-1)
	QuickSort(li, i+1, right)
}

//快速排序-降序
func QuickSortByDESC(li []int, left, right int) {
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
	QuickSortByDESC(li, left, i-1)
	QuickSortByDESC(li, i+1, right)
	return
}
