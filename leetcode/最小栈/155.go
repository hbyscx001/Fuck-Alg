/*
 * @lc app=leetcode.cn id=155 lang=golang
 *
 * [155] 最小栈
 */

// @lc code=start
type MinStack struct {
	data []int
	aux []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}


func (this *MinStack) Push(x int)  {
	if len(this.data) != 0 && x > this.aux[len(this.aux) - 1] {
		this.aux = append(this.aux, this.aux[len(this.aux) - 1])
	} else {
		this.aux = append(this.aux, x)
	}

	this.data = append(this.data, x)
}


func (this *MinStack) Pop()  {
	this.data = this.data[:len(this.data) - 1]
	this.aux = this.aux[:len(this.aux) - 1]
}


func (this *MinStack) Top() int {
	return this.data[len(this.data) - 1]
}


func (this *MinStack) GetMin() int {
	return this.aux[len(this.aux) - 1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end

