package goutils

import (
	"fmt"
	"github.com/zjllib/goutils/logs"
	"github.com/zjllib/goutils/sort"
	"testing"
)

func TestQuickSort(t *testing.T) {
	logs.Info("error")
	logs.Warn("error")
	logs.Error("error")
	arr := []int{1, 10, 7, 9}
	sort.QuickSortByDESC(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
