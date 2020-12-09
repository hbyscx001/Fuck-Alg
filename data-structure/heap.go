package datastrucure

// MinIntHeap is a type of min heap
type MinIntHeap struct {
	Data []int
}

// NewMinIntHeap is constructor of MinIntHeap
func NewMinIntHeap() *MinIntHeap {
	return &MinIntHeap{}
}

// Insert is a method to add a new element to the current heap
func (h *MinIntHeap) Insert(v int) {
	h.Data = append(h.Data, v)
	last := len(h.Data) - 1

	for last/2 >= 0 {
		if h.Data[last/2] > h.Data[last] {
			h.Data[last/2], h.Data[last] = h.Data[last], h.Data[last/2]
			last = last / 2
		} else {
			break
		}
	}
}

// Pop is a method to pop the min element of current heap
func (h *MinIntHeap) Pop() (int, bool) {
	if len(h.Data) < 1 {
		return 0, false
	}

	for child := 1; child < len(h.Data); child = child * 2 {
		if child+1 < len(h.Data) && h.Data[child+1] < h.Data[child] {
			child++
		}

		h.Data[child/2], h.Data[child] = h.Data[child], h.Data[child/2]
	}

	ret := h.Data[len(h.Data)-1]
	h.Data = h.Data[:len(h.Data)-1]
	return ret, true
}