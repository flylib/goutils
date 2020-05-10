package goutils

import (
	"fmt"
	"github.com/Quantumoffices/goutils/sort"
	"testing"
)

func TestQuickSort(t *testing.T) {
	arr := []int{1, 10, 7, 9}
	sort.QuickSortByDESC(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
