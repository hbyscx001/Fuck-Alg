package sort

// QuickSortInt is a function to sort int array
func QuickSortInt(arr []int) {
	if len(arr) < 2 {
		return
	}

	_quickSort(arr, 0, len(arr)-1)
}

func _quickSort(arr []int, left, right int) {
	if left >= right {
		return
	}

	mid := partition(arr, left, right)
	_quickSort(arr, left, mid)
	_quickSort(arr, mid+1, right)
}

// 设置一个哨兵元素，pivot,
// 经过分区函数后，pivot左边的都小于pivot，pivot右边的大于或等与pivot
func partition(arr []int, left, right int) int {
	pivot := left
	left = left + 1

	for left < right {
		if arr[left] < arr[pivot] {
			left++
			continue
		}

		if arr[right] >= arr[pivot] {
			right--
			continue
		}

		swap(arr, left, right)
	}

	// left的值这时候有不同的情况
	// case1, left的值是大于或等于pivot的，这时候需要交换的是left-1和pviot
	// case2, left的值也有可能是小于pivot的，这时候需要交换的仅仅是left和pviot
	if arr[left] >= arr[pivot] {
		swap(arr, pivot, left-1)
	} else {
		swap(arr, pivot, left)
	}

	return left - 1
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
