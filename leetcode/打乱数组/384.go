package leetcode

/*
 * @lc app=leetcode.cn id=384 lang=golang
 *
 * [384] 打乱数组
 */

// @lc code=start

import (
	"math/rand"
	"time"
)


type Solution struct {
	raw []int
	size int
	seed *rand.Rand
}


func Constructor(nums []int) Solution {
	return Solution{
		raw: nums,
		size: len(nums),
		seed: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}


/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	return this.raw
}


/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	newArray := make([]int, this.size)
	copy(newArray, this.raw)
	for i, _ := range newArray {
		randomIndex := this.seed.Intn(this.size)
		newArray[i], newArray[randomIndex] = newArray[randomIndex], newArray[i]
	}
	return newArray
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
// @lc code=end

