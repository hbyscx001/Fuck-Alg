package leetcode

import (
	"fmt"
	"testing"
)


type question88 struct {
	para88
	ans88
}

// para 是参数
// one 代表第一个参数
type para88 struct {
	nums1 []int
	m int
	nums2 []int
	n int
}

// ans 是答案
// one 代表第一个答案
type ans88 struct {
	nums []int
}

func TestProblem88(t *testing.T) {
	qs := []question88{

		{
			para88{[]int{1,2,3, 0, 0, 0}, 3, []int{2, 5, 6}, 3},
			ans88{[]int{1,2, 2, 3, 5, 6}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 88------------------------\n")

	for _, q := range qs {
		expect, p := q.ans88.nums, q.para88
		fmt.Printf("[Input] %v", p)
		merge(p.nums1, p.m, p.nums2, p.n)
		for i, v := range expect {
			if v != p.nums1[i] {
				t.Errorf("Expect: %v Actual: %v", expect, p.nums1)
			}
		}
	}
	fmt.Printf("\n\n\n")
}