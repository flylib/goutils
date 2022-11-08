package goutils

import (
	"fmt"
	"github.com/zjllib/goutils/logs"
	"github.com/zjllib/goutils/sort"
	"testing"
)

type Student struct {
	Name string
}

func TestQuickSort(t *testing.T) {
	logs.Info("error %#v", Student{Name: "张"})
	logs.Warn("error")
	logs.Error("error")
	arr := []int{1, 10, 7, 9}
	sort.QuickSortDESC(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
