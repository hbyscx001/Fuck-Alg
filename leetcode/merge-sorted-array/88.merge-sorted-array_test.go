package leetcode

import (
	"fmt"
	"testing"
)


type question88 struct {
	para4
	ans4
}

// para 是参数
// one 代表第一个参数
type para88 struct {
	nums1 []int
	nums2 []int
}

// ans 是答案
// one 代表第一个答案
type ans88 struct {
	nums []int
}

func TestProblem88(t *testing.T) {
	qs := []question88{

		{
			para88{[]int{1,2,3, 0, 0, 0}, []int{2, 5, 6}},
			ans88{[]int{1,2, 2, 3, 5, 6}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 88------------------------\n")

	for _, q := range qs {
		expect, p := q.ans88.nums, q.para88
		actual := merge(p.nums1, len(p.nums1), p.nums2, len(p.nums2))
		fmt.Printf("【input】:%v       【output】:%v\n", p, actual)
		if actual != expect {
			t.Errorf("Expect: %v", expect)
		}
	}
	fmt.Printf("\n\n\n")
}