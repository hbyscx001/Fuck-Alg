package sort

type CompareFunction func(a, b interface{}) int

func QuickSort(arr []interface{}, compare CompareFunction) {
	if len(arr) < 2 {
		return
	}

	_quickSort(compare, arr, 0, len(arr) - 1)
}

func _quickSort(compare CompareFunction, arr []interface{}, left, right int) {
	if left == right {
		return
	}

	mid := partition(compare, arr, left, right)
	_quickSort(compare, arr, left, mid)
	_quickSort(compare, arr, mid+1, right)
}

func partition(compare CompareFunction, arr []interface{}, left, right int) int {
	pivot := left
	start := left+1
	end := right

	for {
		if start < right && compare(arr[pivot], arr[start]) < 0 {
			start++
		}

		if end > 0 && compare(arr[pivot], arr[end]) > 0 {
			end--
		}

		if start < end {
			swap(arr, start, end)
		} else {
			swap(arr, start, end)
			break
		}
	}

	return end
}

func swap(arr []interface{}, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
