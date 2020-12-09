package leetcode

/*
 * @lc app=leetcode.cn id=295 lang=golang
 *
 * [295] 数据流的中位数
 */

// @lc code=start
type MedianFinder struct {
	data []int
}


/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{}
}


func (this *MedianFinder) AddNum(num int)  {
	start := 0
	end := len(this.data) - 1
	insertPosition := 0

	for start <= end {
		if num < this.data[start] {
			insertPosition = start
			break
		} else if num > this.data[end] {
			insertPosition = end + 1
			break		
		} else {
			mid := (start + end) / 2
			if num < this.data[mid] {
				end = mid - 1
			} else if num > this.data[mid] {
				start = mid + 1
			} else {
				for mid <= end && this.data[mid] == num {
					mid++
				}
				insertPosition = mid
				break
			}
		}
	}

	if insertPosition == len(this.data) {
		this.data = append(this.data, num)
	} else {
		tmp := append([]int{}, this.data[insertPosition:]...)
		this.data = append(this.data[:insertPosition], num)
		this.data = append(this.data, tmp...)
	}
}


func (this *MedianFinder) FindMedian() float64 {
	if odd, div := len(this.data) % 2, len(this.data) / 2; odd == 1 {
		return float64(this.data[div])
	} else {
		return float64(this.data[div] + this.data[div-1]) / 2
	}
}


/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
// @lc code=end

