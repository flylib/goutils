package sort

// 快速排序-沈旭
func partition(list []int, left, right int) int {
	pivot := list[left] //导致 left 位置值为空
	for left < right {
		// >= pivot 指针👈移
		for left < right && pivot <= list[right] {
			right--
		}
		//小于基准的往左放
		list[left] = list[right]
		//left指针值 <= pivot 指针👉移
		for left < right && pivot >= list[left] {
			left++
		}
		//大于基准的往右放
		list[right] = list[left]
	}
	//pivot 填补 left位置的空值
	list[left] = pivot
	return left
}

/*
* 快排升序
 */
func QuickSort(list []int, left, high int) {
	if high > left {
		//位置划分
		pivot := partition(list, left, high)
		//左边部分排序
		QuickSort(list, left, pivot-1)
		//右边排序
		QuickSort(list, pivot+1, high)
	}
}
