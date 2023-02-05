package heap

import "container/heap"

type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() any {
	// pop and remove the last item
	last := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return last
}

func NewMinHeap() *MinHeap {
	return &MinHeap{}
}

func (h *MinHeap) PushVal(x int) {
	heap.Push(h, x)
}

func (h *MinHeap) PopMin() int {
	return heap.Pop(h).(int)
}
