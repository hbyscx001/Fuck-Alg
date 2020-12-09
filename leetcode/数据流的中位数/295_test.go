package leetcode

import (
	"fmt"
	"testing"
)


func Test295(t *testing.T) {
	m := Constructor()

	m.AddNum(0)
	m.AddNum(0)

	fmt.Println(m.FindMedian())
}