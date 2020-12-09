package sort

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	arr := []int{2, 1, 3, 0, 7, -1, -10, 100}

	QuickSort(arr)

	fmt.Println(arr)
}
