package sort

import "testing"

func TestQuickSort(t *testing.T) {
	arr := []int{5, 1, 1, 2, 0, 0}
	QuickSort(arr, 0, len(arr)-1)
	t.Log(arr)
}
