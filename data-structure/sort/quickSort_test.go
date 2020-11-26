package sort

import "testing"

func intCompare(a, b interface{}) int {
	return b.(int) - a.(int)
}

func TestQuickSort(t *testing.T) {
	arr := []int{2,1,3}

	QuickSort(arr, intCompare)
}